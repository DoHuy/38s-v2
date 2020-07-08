package main

import (
	"38s-v2/src/38s/config"
	"38s-v2/src/38s/repositories"
	"38s-v2/src/38s/services"
	"fmt"
	"github.com/go-redis/redis"
	"time"

	//"38s-v2/src/38s/repositories"
	"38s-v2/src/38s/routers"
	"38s-v2/src/pkgs"
	"38s-v2/src/pkgs/clients"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	//"time"

	//"github.com/gin-contrib/sessions"
	//"github.com/go-redis/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	conf, err := config.GetAuthConfig()
	if err != nil {
		panic(err)
	}

	router := gin.New()
	// monitor with newleric
	//router := oneginrelic.InitGinWithNewRelic()
	redisClient, err := clients.NewRedis(&redis.Options{
		Addr:        conf.RedisConfig.Host + ":" + conf.RedisConfig.Port,
		Password:    conf.RedisConfig.Password,
		DB:          conf.RedisConfig.DB,
		MaxRetries:  3,
		IdleTimeout: 5 * time.Minute,
	})
	if err != nil {
		panic(err)
	}

	redisStore := repositories.NewStore(redisClient)

	//dbClient, err := clients.NewDB("mysql", conf.DB.MySQLConnectionString)
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer onClose(redisClient, dbClient)

	//otpRepository := repositories.NewOTPRepository(dbClient)
	//userRepository := repositories.NewUsersRepository(dbClient)

	sessionStore := clients.NewSessionStore(conf.RedisConfig.Host, conf.RedisConfig.Port, conf.RedisConfig.Password, conf.RedisConfig.SessionSecret, sessions.Options{
		MaxAge: conf.RedisConfig.MaxAge,
	})
	router.Use(sessions.Sessions("38s", sessionStore))

	//oauth2Service := services.NewOAuth2Service(conf)
	//sendOTPService := services.NewSendOTP(conf)
	//otpService := services.NewOTPService(conf, redisStore, sendOTPService, otpRepository)
	//userService := services.NewUserService(conf, userRepository)
	sessionService := services.NewSessionService()

	server := routers.NewServer(router, redisStore, sessionService, conf)

	pkgs.HandleSigterm(func() {
		if err := server.Stop(); err != nil {
			fmt.Println(nil, "Failed to Stop API server", err)
		}
	})

	if err := server.Start(); err != nil {
		if err == http.ErrServerClosed {
			fmt.Println(nil, "Shutdown API server")
		} else {
			fmt.Println(nil, "Failed to Shutdown API server", err)
		}
	}
}

//func onClose(clients ...io.Closer) {
//	for _, client := range clients {
//		err := client.Close()
//		if err != nil {
//			log.Fatal(nil, "Failed to close client", err)
//		}
//	}
//}

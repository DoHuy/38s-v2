package clients

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

func NewSessionStore(host, port, password, secret string, options sessions.Options) redis.Store {
	sessionStore, err := redis.NewStore(3, "tcp", host+":"+port, password, []byte(secret))
	if err != nil {
		panic(err)
	}
	sessionStore.Options(options)
	return sessionStore
}

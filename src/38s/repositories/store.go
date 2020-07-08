package repositories

import (
	"time"
	"github.com/go-redis/redis"
)

// Store interface
type Store interface {
	Set(key string, value string, expireTime int64) error
	SetWithTimes(key string, expired uint) (uint, error)
	Get(key string) (string, error)
	GetInt(key string) (int64, error)
	Increase(key string, expireTime int64) (int64, error)
	Exists(key string) (bool, error)
	Del(key string) (int64, error)
	TTL(key string) (time.Duration, error)
	HSet(key string, field string, value string, expired uint) (bool, error)
	HSetInt(key string, field string, value int64, expired int64) (bool, error)
	HGet(key string, field string) (string, error)
	HGetInt(key string, field string) (int64, error)
	HDel(key string, field string) (bool, error)
	Expire(key string, expire int64) (bool, error)
	Ping() error
}

type store struct {
	client *redis.Client
}

func NewStore(client *redis.Client) Store {
	return &store{
		client:client}
}

func (s *store) Set(key string, value string, expired int64) error {
	exist, err := s.client.Get(key).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	if exist == "1" {
		return nil
	}
	_, err = s.client.Set(key, value, time.Duration(expired)*time.Second).Result()
	return err
}

func (s *store) SetWithTimes(key string, expired uint) (uint, error) {
	var setExpired = false
	_, err := s.client.Get(key).Result()
	if err == redis.Nil {
		setExpired = true
	}
	result, err := s.client.Incr(key).Result()
	if err != nil {
		return 0, err
	}
	if setExpired == true {
		_, err := s.client.Expire(key, time.Duration(expired)*time.Second).Result()
		if err != nil {
			return 0, err
		}
	}
	return uint(result), nil
}

func (s *store) Get(key string) (string, error) {
	str, err := s.client.Get(key).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return str, nil
}

func (s *store) GetInt(key string) (int64, error) {
	value, err := s.client.Get(key).Int64()
	if err != nil && err != redis.Nil {
		return -1, err
	}
	return value, nil
}

func (s *store) Increase(key string, expireTime int64) (int64, error) {
	value, err := s.client.Incr(key).Result()
	if err != nil {
		return -1, err
	}
	if expireTime > 0 {
		s.client.Expire(key, time.Duration(expireTime)*time.Second)
	}
	return value, nil
}

func (s *store) Exists(key string) (bool, error) {
	value, err := s.client.Exists(key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}

	if value == 1 {
		return true, nil
	}

	return false, nil
}

func (s *store) Del(key string) (int64, error) {
	return s.client.Del(key).Result()
}

func (s *store) TTL(key string) (time.Duration, error) {
	result, err := s.client.TTL(key).Result()
	if err != nil && err != redis.Nil {
		return -1, err
	}
	ttl := result / 1000000000
	return ttl, nil
}

func (s *store) HSet(key string, field string, value string, expired uint) (bool, error) {
	result, err := s.client.HSet(key, field, value).Result()
	if err != nil {
		return false, err
	}
	if expired > 0 {
		_, err = s.client.Expire(key, time.Duration(expired)*time.Second).Result()
		if err != nil {
			return false, err
		}
	} else {
		_, err := s.client.Expire(key, time.Duration(86400)*time.Second).Result()
		if err != nil {
			return false, err
		}
	}
	return result, nil
}

func (s *store) HSetInt(key string, field string, value int64, expired int64) (bool, error) {
	result, err := s.client.HSet(key, field, value).Result()
	if err != nil {
		return false, err
	}

	_, err = s.client.Expire(key, time.Duration(expired)*time.Second).Result()
	if err != nil {
		return false, err
	}

	return result, nil
}

func (s *store) HGet(key string, field string) (string, error) {
	value, err := s.client.HGet(key, field).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}
		return "", err
	}
	return value, nil
}

func (s *store) HGetInt(key string, field string) (int64, error) {
	value, err := s.client.HGet(key, field).Int64()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}

	return value, nil
}

func (s *store) HDel(key string, field string) (bool, error) {
	_, err := s.client.HDel(key, field).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *store) Expire(key string, expire int64) (bool, error) {
	return s.client.Expire(key, time.Duration(expire)).Result()
}

func (s *store) Ping() error {
	_, err := s.client.Ping().Result()
	return err
}
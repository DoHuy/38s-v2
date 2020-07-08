package services

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionService interface {
	Start(ctx *gin.Context)
	Get(key interface{}) interface{}
	Set(key interface{}, val interface{})
	Delete(key interface{})
	AddFlash(value interface{}, vars ...string)
	Flashes(vars ...string) []interface{}
	Save() error
}

type SessionServiceImpl struct {
	session sessions.Session
}

func (s *SessionServiceImpl) Start(ctx *gin.Context) {
	s.session = sessions.Default(ctx)
}

func (s *SessionServiceImpl) Save() error {
	return s.session.Save()
}

func (s *SessionServiceImpl) Get(key interface{}) interface{} {
	return s.session.Get(key)
}

func (s *SessionServiceImpl) Set(key interface{}, val interface{}) {
	s.session.Set(key, val)
}

func (s *SessionServiceImpl) Delete(key interface{}) {
	s.session.Delete(key)
}

func (s *SessionServiceImpl) AddFlash(value interface{}, vars ...string) {
	s.session.AddFlash(value, vars...)
}

func (s *SessionServiceImpl) Flashes(vars ...string) []interface{} {
	return s.session.Flashes(vars...)
}

func NewSessionService() SessionService {
	return &SessionServiceImpl{}
}

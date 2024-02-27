package server

import "github.com/gin-gonic/gin"

type Api struct {
	listenAddr string
}

func New(listenAddr string) *Api {
	return &Api{
		listenAddr: listenAddr,
	}
}

func (a *Api) Start() error {

	r := gin.Default()

	r.POST("/mock/login", a.login)

	return r.Run(a.listenAddr)
}

package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Qora struct {
}

func New() *Qora {
	return &Qora{}
}

func (qora *Qora) Init() {

}

func (qora *Qora) Start() {
	// apply default Gin service
	router := gin.Default()
	// apply Gin logger & recovery middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// create router group for nova
	qoraService := router.Group("qora/v1")
	{
		qoraService.GET("/test", func(c *gin.Context) { c.String(http.StatusOK, "hello Qora\n") })
	}
	// create http server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}

func (qora *Qora) Stop() {

}

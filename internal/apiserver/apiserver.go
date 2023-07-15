package apiserver

import (
	"log"
	"net/http"

	"github.com/avg73/realworld-go-api/internal/config"
	"github.com/avg73/realworld-go-api/internal/logger"
	"github.com/avg73/realworld-go-api/internal/router"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *config.Config
	logger *logrus.Logger
	router *mux.Router
}

func New(c *config.Config) *APIServer {
	return &APIServer{
		config: c,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *APIServer) Start() {
	err := logger.ConfigureLogger(a.logger, a.config.LogLevel)
	if err != nil {
		log.Fatal("logger configure error: ", err)
	}

	err = router.ConfigureRouter(a.router)
	if err != nil {
		log.Fatal("router configure error: ", err)
	}

	a.logger.Info("starting api server")
	http.ListenAndServe(a.config.BindAddr, a.router)
}

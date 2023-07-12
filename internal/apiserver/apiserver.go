package apiserver

import (
	"log"

	"github.com/avg73/realworld-go-api/internal/config"
	"github.com/avg73/realworld-go-api/internal/logger"
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
	}
}

func (s *APIServer) Start() {
	err := logger.ConfigureLogger(s.logger, s.config.LogLevel)
	if err != nil {
		log.Fatal("logger configure error: ", err)
	}
	s.logger.Info("starting api server")
}

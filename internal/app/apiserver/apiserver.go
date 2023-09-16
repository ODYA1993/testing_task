package apiserver

import (
	"github.com/DmitryOdintsov/rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func NewAPIServer(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configLogger(); err != nil {
		return err
	}

	s.configRouter()

	if err := s.configStore(); err != nil {
		return err
	}

	s.logger.Info("Starting API server...")

	return http.ListenAndServe(s.config.BinAddr, s.router)
}

func (s *APIServer) configLogger() error {
	lever, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(lever)

	return nil
}

func (s *APIServer) configRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configStore() error {
	st := store.NewStore(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st

	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "Hello")
	}
}

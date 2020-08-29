package apiserver

import (
	"github.com/niqitosiq/BDase/internal/app/chain"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
	"io"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	};
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter();

	s.logger.Info("Server Started")


	return http.ListenAndServe(s.config.BindAddr, s.router);
}


func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/create", s.handleCreate())
	s.router.HandleFunc("/newBlock", s.handleNewBlock())
}

func (s *APIServer) handleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		newChain := chain.NewChain("chain")
		io.WriteString(w, newChain.Name)
	}
}

func (s *APIServer) handleNewBlock() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		currentChain := chain.AppendBlock("chain", "Новый контент")

		var contents string = ""
		for _, block := range currentChain.Blocks {
			contents = contents + block.Content
		}

		io.WriteString(w, contents)
	}
}
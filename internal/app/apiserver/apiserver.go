package apiserver

import (
	"github.com/niqitosiq/BDase/internal/app/chain"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"io"
	"io/ioutil"
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


// ChainCreate ...
type ChainCreate struct {
	Name string `json:"name"`
}
func (s *APIServer) handleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content :=  ChainCreate{}

		body, bodyErr := ioutil.ReadAll(r.Body)
		if bodyErr != nil {
			http.Error(w, bodyErr.Error(), http.StatusBadRequest)
			return
		}

		jsonErr := json.Unmarshal(body, &content)
		if jsonErr != nil {
			http.Error(w, jsonErr.Error(), http.StatusBadRequest)
			return
		}

		newChain := chain.NewChain(content.Name)
		
		s.logger.Info("New chain inited: ", content.Name)

		io.WriteString(w, newChain.Name)
	}
}

// NewBlockContent ...
type NewBlockContent struct {
	Chain string `json:"chain"`
	Content string `json:"content"`
}
func (s *APIServer) handleNewBlock() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content :=  NewBlockContent{}

		body, bodyErr := ioutil.ReadAll(r.Body)
		if bodyErr != nil {
			http.Error(w, bodyErr.Error(), http.StatusBadRequest)
			return
		}

		jsonErr := json.Unmarshal(body, &content)
		if jsonErr != nil {
			http.Error(w, jsonErr.Error(), http.StatusBadRequest)
			return
		}

		currentChain := chain.AppendBlock(content.Chain, content.Content)

		s.logger.Info("New data in '", content.Chain,"': ",content.Content)

		var contents string = ""
		for _, block := range currentChain.Blocks {
			contents = contents + block.Content
		}

		io.WriteString(w, contents)
	}
}
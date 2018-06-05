package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/psucodervn/iir/pkg/parsers"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

func mainServer(cmd *cobra.Command, args []string) error {
	var server Server
	if serverType == "json" {
		server = new(JSONServer)
	} else {
		server = new(HTMLServer)
	}
	http.HandleFunc("/", server.Handler)
	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, nil)
	return err
}

// Server interface
type Server interface {
	Handler(http.ResponseWriter, *http.Request)
}

// HTMLServer is basic implement of Server interface
// It receive html from web extension
type HTMLServer struct {
}

var (
	// ErrWrongHTMLFormat wrong html format
	ErrWrongHTMLFormat = errors.New("wrong html format")
	// ErrParserNotSupported parser not supported
	ErrParserNotSupported = errors.New("parser not supported")
)

func (s *HTMLServer) parseHTML(html string) error {
	nextLineIndex := strings.IndexAny(html, "\r\n")
	if nextLineIndex < 0 {
		return ErrWrongHTMLFormat
	}
	site := html[:nextLineIndex]
	parser := parsers.NewParser(site)
	if parser == nil {
		return ErrParserNotSupported
	}
	task, err := parser.ParseTaskFromHTML(html[nextLineIndex:])
	log.Info().Err(err).Msgf("%v", task)
	return nil
}

// Handler is the handler function of server
func (s *HTMLServer) Handler(writer http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Error().Err(err).Msg("ReadAll body")
		return
	}

	html := string(data)
	log.Debug().Msg(html)
	if s.parseHTML(html) != nil {
		writer.WriteHeader(500)
	} else {
		writer.WriteHeader(200)
	}
}

// JSONServer is basic implement of Server interface
// It receive json from web extension
type JSONServer struct {
}

func (s *JSONServer) parseJSON(data []byte) error {
	var task parsers.Task
	if err := json.Unmarshal(data, &task); err != nil {
		log.Error().Err(err).Msg("AddTask")
		log.Info().Msgf("%s", string(data))
		return err
	}

	if err := parsers.AddTask(task); err != nil {
		log.Error().Err(err).Msg("AddTask")
		log.Info().Msgf("%#v", task)
	}

	log.Info().Msgf("Parse task %v done.", task.Name)
	return nil
}

// Handler is the handler function of server
func (s *JSONServer) Handler(writer http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Error().Err(err).Msg("ReadAll body")
		return
	}

	if s.parseJSON(data) != nil {
		writer.WriteHeader(500)
	} else {
		writer.WriteHeader(200)
	}
}

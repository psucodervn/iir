package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

func mainServer(cmd *cobra.Command, args []string) error {
	server := &BasicServer{}
	http.HandleFunc("/", server.Handler)
	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, nil)
	return err
}

// Server interface
type Server interface {
	Handler(http.ResponseWriter, *http.Request)
}

// BasicServer is basic implement of Server interface
type BasicServer struct {
}

// Handler is the handler function of server
func (s *BasicServer) Handler(writer http.ResponseWriter, request *http.Request) {
	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Error().Err(err).Msg("ReadAll body")
		return
	}
	fmt.Println(string(data))
}

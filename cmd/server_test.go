package cmd

import (
	"net/http"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestBasicServerImplementsServer(t *testing.T) {
	as := assert.New(t)

	as.Implements((*Server)(nil), new(BasicServer))
}

func Test_mainServer(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := mainServer(tt.args.cmd, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("mainServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBasicServer_Handler(t *testing.T) {
	type args struct {
		writer  http.ResponseWriter
		request *http.Request
	}
	var tests []struct {
		name string
		s    *BasicServer
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &BasicServer{}
			s.Handler(tt.args.writer, tt.args.request)
		})
	}
}

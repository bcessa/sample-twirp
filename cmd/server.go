package cmd

import (
	"fmt"
	"github.com/bcessa/sample-twirp/proto"
	"github.com/bcessa/sample-twirp/rpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start sample server",
	RunE:  runServer,
}

func init() {
	var port int
	serverCmd.Flags().IntVar(&port, "port", 9000, "tcp port to use")
	viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))
	rootCmd.AddCommand(serverCmd)
}

func runServer(_ *cobra.Command, _ []string) error {
	endpoint := fmt.Sprintf(":%d", viper.GetInt("server.port"))
	log.Printf("starting server on %s", endpoint)
	handler := sample.NewBusinessCaseServer(&rpc.SampleServer{}, nil)
	return http.ListenAndServe(endpoint, handler)
}

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

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Start sample client",
	RunE:  runClient,
}

func init() {
	var (
		port    int
		useJSON bool
	)
	clientCmd.Flags().IntVar(&port, "port", 9000, "tcp port to use")
	clientCmd.Flags().BoolVar(&useJSON, "json", false, "use JSON client")
	viper.BindPFlag("client.port", clientCmd.Flags().Lookup("port"))
	viper.BindPFlag("client.json", clientCmd.Flags().Lookup("json"))
	rootCmd.AddCommand(clientCmd)
}

func runClient(_ *cobra.Command, _ []string) error {
	endpoint := fmt.Sprintf("http://localhost:%d", viper.GetInt("client.port"))
	log.Printf("starting client on %s", endpoint)

	// Get client
	var client sample.BusinessCase
	if viper.GetBool("client.json") {
		log.Println("using JSON")
		client = sample.NewBusinessCaseJSONClient(endpoint, &http.Client{})
	} else {
		log.Println("using Protocol Buffers")
		client = sample.NewBusinessCaseProtobufClient(endpoint, &http.Client{})
	}

	// Start console
	log.Println("starting interactive commands console")
	cli := rpc.NewConsole(client, "\033[33m»\033[0m ")
	defer cli.Close()
	return cli.Start()
}

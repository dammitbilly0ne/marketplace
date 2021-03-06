package cmd

import (
	"fmt"
	"github.com/dammitbilly0ne/marketplace/internal/handlers"
	marketplaceapialpha "github.com/dammitbilly0ne/marketplace/protos"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
)

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "Run the gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("starting server")
		port := 5000
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()

		handler, err := handlers.NewAlpha(&handlers.AlphaConfig{})
		if err != nil {
			log.Fatal("error returned fro handlers.NewAlpha. ", err)
		}

		marketplaceapialpha.RegisterMarketplaceAPIServer(s,handler)

		log.Printf("server listening at %v", lis.Addr())

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}

func init() {
	rootCommand.AddCommand(serverCommand)
}

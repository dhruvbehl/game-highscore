package main

import (
	"flag"

	grpcBackend "github.com/dhruvbehl/game-highscore/internal/server/grpc"
	"github.com/rs/zerolog/log"
)

/*
Usage: go run main.go -address=:<port>
Summary:
1. Get new `server` from the given `address`
2. ListenAndServe using the `server`
*/
func main() {
	address := flag.String("address", ":9001", "address to connect to game-highscore service")
	flag.Parse()
	server := grpcBackend.NewServer(*address)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize gRPC server for service game-highscore ")
	}
}

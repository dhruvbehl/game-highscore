package main

import (
	"flag"
	grpcBackend "github.com/dhruvbehl/game-highscore/internal/server/grpc"
	"github.com/rs/zerolog/log"
)

func main() {
	address := flag.String("address", ":9001", "address to connect to game-highscore service")
	flag.Parse()
	srv := grpcBackend.NewServer(*address)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize gRPC server for service game-highscore ")
	}
}
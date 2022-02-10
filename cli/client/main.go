package main

import (
	"flag"
	"time"

	pbhighscore "github.com/dhruvbehl/game-apis/game-highscore/v1"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

/*
Usage: go run main.go -address=localhost:<port>
Summary:
1. Creates a dial connection to given `address` to get a `connection`
2. Use `connection` to get `client`
3. Call API using `client`
*/
func main() {
	address := flag.String("address", "localhost:9001", "address to connect to game-highscore service")
	flag.Parse()

	connection, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to dial game-highscore gRPC service")
	}

	// Connection clean up that runs at the end
	defer func(){
		err := connection.Close()
		if err != nil {
			log.Fatal().Err(err).Str("address", *address).Msg("failed to close the connection")
		}
	}()

	// Cancel context if time exceeds 10s
	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client := pbhighscore.NewGameClient(connection)
	if client == nil {
		log.Info().Msg("Client nil")
	}

	request, err := client.GetHighScore(timeoutCtx, &pbhighscore.GetHighScoreRequest{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get a response")
	}

	if request != nil {
		log.Info().Interface("highscore", request.GetHighScore()).Msg("Highscore received from game-highscore service")
	} else {
		log.Error().Msg("Couldn't get highscore")
	}
}
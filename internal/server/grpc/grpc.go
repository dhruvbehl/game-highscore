package grpc

import (
	"context"
	"math"

	"net"

	pbhighscore "github.com/dhruvbehl/game-apis/game-highscore/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

var HighScore = math.MaxFloat64

func NewServer(address string) *Grpc {
	return &Grpc{address: address}
}

func (g *Grpc) SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest) (*pbhighscore.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore in game-highscore")
	HighScore = input.HighScore
	return &pbhighscore.SetHighScoreResponse{
		Set: true,
	}, nil
}

func (g *Grpc) GetHighScore(ctx context.Context, input *pbhighscore.GetHighScoreRequest) (*pbhighscore.GetHighScoreResponse, error) {
	log.Info().Msg("GetHighScore in game-highscore")
	return &pbhighscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	listener, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "Failed to initialize tcp port")
	}

	serverOpts := []grpc.ServerOption{}
	g.srv = grpc.NewServer(serverOpts...)

	pbhighscore.RegisterGameServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("Initializing gRPC server for game-highscore service")

	err = g.srv.Serve(listener)
	if err != nil {
		return errors.Wrap(err, "Failed to initialize gRPC server for game-highscore service")
	}
	return nil
}

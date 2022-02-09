package grpc

import (
	"context"
	"math"

	pbhighscore "github.com/dhruvbehl/game-apis/game-highscore/v1"
	"google.golang.org/grpc"
	"github.com/rs/zerolog/log"
)

type Grpc struct {
	address string
	srv		*grpc.Server
}

var HighScore = math.MaxFloat64

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
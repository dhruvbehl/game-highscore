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

func (grpc *Grpc) SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest) (*pbhighscore.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore in game-highscore")
	HighScore = input.HighScore
	return &pbhighscore.SetHighScoreResponse{
		Set: true,
	}, nil
}
package server

import (
	"context"
	"math/rand"
	"time"
)

func randomSleep(ctx context.Context, maxSeconds int) int {

	seconds := rand.Intn(maxSeconds)

	duration := time.Duration(seconds) * time.Second

	select {
	case <-time.After(duration):
		//sleep completed
		return seconds
	case <-ctx.Done():
		//context was cancelled
		return seconds
	}

}

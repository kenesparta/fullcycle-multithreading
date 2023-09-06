package requests

import (
	"context"

	"github.com/kensparta/fullcycle-multithreading/dto"
)

type BuildInterface[T any] interface {
	Execute(ctx context.Context, result chan dto.ResultChannel[T])
}

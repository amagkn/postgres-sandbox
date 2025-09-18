package ver1

import (
	"github.com/amagkn/postgres-sandbox/internal/air/usecase"
)

type Handlers struct {
	uc *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{uc: uc}
}

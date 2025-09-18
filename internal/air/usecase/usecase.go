package usecase

type Postgres interface{}

type UseCase struct {
	postgres Postgres
}

func New(p Postgres) *UseCase {
	return &UseCase{
		postgres: p,
	}
}

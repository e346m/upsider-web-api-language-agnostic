package usecases

import "go.opentelemetry.io/otel/trace"

type Usecase struct {
	repo   RepositoryKeeper
	tracer trace.Tracer
	auther Auther
}

func NewUsecase(repo RepositoryKeeper, tracer trace.Tracer, auther Auther) *Usecase {
	return &Usecase{
		repo:   repo,
		tracer: tracer,
		auther: auther,
	}
}

package usecases

import "go.opentelemetry.io/otel/trace"

type Usecase struct {
	repo   RepositoryKeeper
	tracer trace.Tracer
}

func NewUsecase(repo RepositoryKeeper, tracer trace.Tracer) *Usecase {
	return &Usecase{
		repo:   repo,
		tracer: tracer,
	}
}

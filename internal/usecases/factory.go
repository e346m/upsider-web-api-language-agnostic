package usecases

type Usecase struct {
	repo RepositoryKeeper
}

func NewUsecase(repo RepositoryKeeper) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

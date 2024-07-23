package usecase

type (
	LoginUseCaseInput struct {
		Email    string
		Password string
	}

	LoginUseCaseOutput struct {
		UserID string
		Email  string
		Token  string
	}
)

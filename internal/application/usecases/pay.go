package usecases

type PayUseCase interface{}

type payUseCase struct{}

func NewPayUseCase() PayUseCase {
	return &payUseCase{}
}

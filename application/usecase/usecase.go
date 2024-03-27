package usecase

type UseCase[t, y any] interface {
	Execute(input t) y
}

package service

//Random contract for random logic
type Random interface {
	Get() (int, error)
}

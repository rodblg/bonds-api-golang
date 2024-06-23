package bondApi

type Storage interface {
	InsertNewData(Bond) error
}

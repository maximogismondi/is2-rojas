package database

type Database[T any] interface {
	Create(string, T) (T, error)
}

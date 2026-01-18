package interfaces

type Cache[K comparable, V any] interface {
	Get(key K) (V, bool)
	Set(Key K, value V)
	Delete(key K)
	Has(key K) bool
	Clear()
	Size() int
	GetAll() map[K]V
	Keys() []K
}

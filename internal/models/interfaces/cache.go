package interfaces

type Cache interface {
	Add(key interface{}, value interface{})
	Get(key interface{}) (interface{}, bool)
	Remove(key interface{}) bool
	GetAll() interface{}
	GetAllAddresses() []string
	HasAny() bool
	Count() int
	ForEach(fn func(key interface{}, value interface{}))
}

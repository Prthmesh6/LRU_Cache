package Cache

type Cache[T1 comparable, T2 comparable] interface {
	Get(element T1) (T2, error)
	Set(key T1, element T2)
	// GetCapacity() int
}

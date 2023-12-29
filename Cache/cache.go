package Cache

import "context"

type Cache[T1 comparable, T2 comparable] interface {
	
	Get(ctx context.Context, element T1) (T2, error)
	Set(ctx context.Context, key T1, element T2) error
	GetCapacity() (int, error)
}

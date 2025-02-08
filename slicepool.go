package pools

import (
	"sort"
	"sync"
)

type SlicePool struct {
	sizes    []int
	pools    []sync.Pool
	min, max int
	fn       SizeFn
}

func NewSlicePoolDefault(fn SizeFn) (*SlicePool, error) {
	return NewSlicePool(minBufferSize, maxBufferSize, growFactor, fn)
}

func NewSlicePool(min, max, factor int, fn SizeFn) (*SlicePool, error) {
	sizes := make([]int, 0)
	for size := min; size <= max; size *= factor {
		sizes = append(sizes, size)
	}
	pools := make([]sync.Pool, len(sizes))
	for idx, size := range sizes {
		pools[idx].New = fn(size)
	}
	return &SlicePool{
		sizes: sizes,
		pools: pools,
		min: min,
		max: max,
		fn: fn,
	}, nil
}

type SizeFn func(size int) func() interface{}

func (p *SlicePool) Alloc(size int) interface{} {
	i := sort.SearchInts(p.sizes, size)
	if i < len(p.sizes) {
		return p.pools[i].Get()
	}
	return p.fn(size)()
}

func (p *SlicePool) Free(size int, value interface{}) {
	i := sort.SearchInts(p.sizes, size)
	if i < len(p.sizes) {
		p.pools[i].Put(value)
	}
}


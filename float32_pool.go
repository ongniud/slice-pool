package pools

type Float32sPool struct {
	*SlicePool
}

func NewFloat32PoolDefault() (*SlicePool, error) {
	return NewSlicePool(minBufferSize, maxBufferSize, growFactor, allocFloat32s)
}

func NewFloat32Pool(min, max, factor int) (*Float32sPool, error) {
	p, err := NewSlicePool(min, max, factor, allocFloat32s)
	if err != nil {
		return nil, err
	}
	return &Float32sPool{
		SlicePool: p,
	}, nil
}

func allocFloat32s(size int) func() interface{} {
	return func() interface{} {
		return make([]float32, 0, size)
	}
}

func (p *Float32sPool) Get(size int) []float32 {
	if size == 0 {
		return nil
	}
	data, ok := p.Alloc(size).([]float32)
	if !ok {
		return nil
	}
	return data[:0]
}

func (p *Float32sPool) Put(v []float32) {
	if v == nil {
		return
	}
	size := cap(v)
	p.Free(size, v)
}

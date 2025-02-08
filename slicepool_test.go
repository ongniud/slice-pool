package pools

import (
	"testing"
)

func TestSlicePool(t *testing.T) {
	p, err := NewSlicePool(minBufferSize, maxBufferSize, growFactor, allocFloat32s)
	if err != nil {
		panic(err)
	}

	size := 10
	b := p.Alloc(size)
	defer p.Free(size, b)

}

func BenchmarkAcquireRawObjs(b *testing.B)  {
	size := 10000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			a := allocFloat32s(size)()
			b := a.([]float32)
			_ = b
		}
	}
}

func BenchmarkAcquirePooledObjs(b *testing.B)  {
	p, err := NewSlicePool(minBufferSize, maxBufferSize, growFactor, allocFloat32s)
	if err != nil {
		panic(err)
	}

	size := 10000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			acquireSlice(p, size)
		}
	}
}

func acquireSlice(p *SlicePool, size int) {
	b := p.Alloc(size)
	p.Free(size, b)
}

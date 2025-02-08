package pools

import (
	"fmt"
	"testing"
)

// go test slicepool_test.go -bench=. -benchmem
func TestFloat32PoolDefault(t *testing.T) {
	p, err := NewFloat32Pool(minBufferSize, maxBufferSize, growFactor)
	if err != nil {
		panic(err)
	}

	b := p.Get(0)
	defer p.Put(b)
	fmt.Println("len(b)", len(b), "cap(b)", cap(b))
}

func TestFloat32Pool(t *testing.T) {
	p, err := NewFloat32Pool(100, 500, 2)
	if err != nil {
		panic(err)
	}
	b := p.Get(500)
	defer p.Put(b)
	fmt.Println("len(b)", len(b), "cap(b)", cap(b))
}

func BenchmarkFloat32PoolAcquireRawObjs(b *testing.B)  {
	size := 10000
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			a := allocFloat32s(size)()
			b := a.([]float32)
			_ = b
		}
	}
}

func BenchmarkFloat32PoolAcquirePooledObjs(b *testing.B)  {
	p, err := NewFloat32Pool(100, 100000, 2)
	if err != nil {
		panic(err)
	}

	size := 10000
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			acquireFloat32s(p, size)
		}
	}
}

func acquireFloat32s(p *Float32sPool, size int) {
	b := p.Get(size)
	p.Put(b)
}

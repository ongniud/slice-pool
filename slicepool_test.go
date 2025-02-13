package pool

import (
	"reflect"
	"testing"
)

func TestNewSlicePoolDefault(t *testing.T) {
	pool := NewSlicePoolDefault[int]()
	if pool == nil {
		t.Errorf("NewSlicePoolDefault returned nil")
	}
	if len(pool.sizes) == 0 {
		t.Errorf("Pool sizes should not be empty")
	}
	if len(pool.pools) == 0 {
		t.Errorf("Pool pools should not be empty")
	}
}

func TestNewSlicePool(t *testing.T) {
	min := 8
	max := 64
	factor := 2
	pool := NewSlicePool[int](min, max, factor)
	if pool == nil {
		t.Errorf("NewSlicePool returned nil")
	}
	if len(pool.sizes) == 0 {
		t.Errorf("Pool sizes should not be empty")
	}
	if len(pool.pools) == 0 {
		t.Errorf("Pool pools should not be empty")
	}
	expectedSizes := []int{8, 16, 32, 64}
	if !reflect.DeepEqual(pool.sizes, expectedSizes) {
		t.Errorf("Expected sizes %v, but got %v", expectedSizes, pool.sizes)
	}
}

func TestAlloc(t *testing.T) {
	pool := NewSlicePoolDefault[int]()
	size := 32
	slice := pool.Alloc(size)
	if cap(slice) < size {
		t.Errorf("Allocated slice capacity %d is less than requested size %d", cap(slice), size)
	}
}

func TestFree(t *testing.T) {
	pool := NewSlicePoolDefault[int]()
	size := 32
	slice := pool.Alloc(size)
	pool.Free(slice)
	newSlice := pool.Alloc(size)
	if &newSlice[0] != &slice[0] {
		t.Errorf("Slice was not reused after being freed")
	}
}

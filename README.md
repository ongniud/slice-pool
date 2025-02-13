# SlicePool

`SlicePool` is a package in Go that provides a pool of slices of varying sizes for efficient memory management. This package aims to reduce memory allocation overhead by reusing pre-allocated memory slices.

## Features

- **Slice Pooling**: Manages a pool of slices of different sizes to optimize memory allocation and deallocation.
- **Dynamic Sizing**: Allows for dynamic sizing of memory slices based on the specified minimum, maximum, and growth factor.
- **Concurrency Safe**: Utilizes `sync.Pool` to ensure safe concurrent access to the pooled slices.

## Usage

### Initialization

To create a new `SlicePool` with default settings, use the `NewSlicePoolDefault` function.

```go
slicePool, err := NewSlicePoolDefault(sizeFunc)
```

To create a custom `SlicePool` with specific minimum, maximum, and growth factor settings, use the `NewSlicePool` function.

```go
slicePool, err := NewSlicePool(minSize, maxSize, growthFactor, sizeFunc)
```

### Allocation and Deallocation

To allocate a slice of a specific size from the pool, use the `Alloc` method.

```go
size := 100
slice := slicePool.Alloc(size).([]int)
```

To release a previously allocated slice back to the pool, use the `Free` method.

```go
slicePool.Free(size, slice)
```

### Size Function

The `SizeFn` type is used to define a function that returns a new instance of a slice of a specified size.

```go
type SizeFn func(size int) func() interface{}
```

## Example

```go
slicePool, err := NewSlicePoolDefault(func(size int) func() interface{} {
    return func() interface{} {
        return make([]int, size)
    }
})

size := 50
slice := slicePool.Alloc(size).([]int)
defer slicePool.Free(size, slice)
```

## License

This package is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

If you encounter any issues, have suggestions, or would like to contribute, feel free to create a pull request or issue on GitHub.

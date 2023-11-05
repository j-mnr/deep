package deep_test

import (
	"reflect"
	"testing"

	"github.com/j-mnr/deep"
	"github.com/stretchr/testify/assert"
)

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type integer interface {
	signed | unsigned
}

type float interface {
	~float32 | ~float64
}

type cmplex interface {
	~complex64 | ~complex128
}

type ordered interface {
	integer | float | cmplex | ~string
}

func TestCopy(t *testing.T) {
	t.Parallel()
	for name, tc := range map[string]struct {
		data any
	}{
		"bool true": {
			data: true,
		},
		"bool false": {
			data: false,
		},

		"int": {
			data: int(8),
		},
		"int8": {
			data: int8(8),
		},
		"int16": {
			data: int16(8),
		},
		"int32": {
			data: int32(8),
		},
		"int64": {
			data: int64(8),
		},

		"uint": {
			data: uint(8),
		},
		"uint8": {
			data: uint8(8),
		},
		"uint16": {
			data: uint16(8),
		},
		"uint32": {
			data: uint32(8),
		},
		"uint64": {
			data: uint64(8),
		},

		"float32": {
			data: float32(8),
		},
		"float64": {
			data: float64(8),
		},

		"complex64": {
			data: complex64(8),
		},
		"complex128": {
			data: complex128(8),
		},

		"uintptr": {
			data: uintptr(8),
		},

		"string": {
			data: "beef",
		},

		"pointer to bool true": {
			data: toPtr(true),
		},
		"pointer to bool false": {
			data: toPtr(false),
		},

		"pointer to int": {
			data: toPtr(int(8)),
		},
		"pointer to int8": {
			data: toPtr(int8(8)),
		},
		"pointer to int16": {
			data: toPtr(int16(8)),
		},
		"pointer to int32": {
			data: toPtr(int32(8)),
		},
		"pointer to int64": {
			data: toPtr(int64(8)),
		},

		"pointer to uint": {
			data: toPtr(uint(8)),
		},
		"pointer to uint8": {
			data: toPtr(uint8(8)),
		},
		"pointer to uint16": {
			data: toPtr(uint16(8)),
		},
		"pointer to uint32": {
			data: toPtr(uint32(8)),
		},
		"pointer to uint64": {
			data: toPtr(uint64(8)),
		},

		"pointer to float32": {
			data: toPtr(float32(8)),
		},
		"pointer to float64": {
			data: toPtr(float64(8)),
		},

		"pointer to complex64": {
			data: toPtr(complex64(8)),
		},
		"pointer to complex128": {
			data: toPtr(complex128(8)),
		},

		"pointer to uintptr": {
			data: toPtr(uintptr(8)),
		},

		"pointer to string": {
			data: toPtr("beef"),
		},

		// "unsafe pointer": {
		// 	data: unsafe.Pointer(toPtr(8)),
		// },

		// Array
		// Chan
		// Func
		// Interface
		// Map
		// Slice
		// Struct
		// UnsafePointer
	} {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			switch reflect.TypeOf(tc.data).Kind() {
			// TODO(jay): These cases.
			// case reflect.Array:
			// case reflect.Chan:
			// case reflect.Func:
			// case reflect.Interface:
			// case reflect.Map:
			case reflect.Slice:
			case
				reflect.Pointer,
				reflect.UnsafePointer:
				got := deep.Copy(tc.data)
				var changed any
				var concrete any
				switch v := tc.data.(type) {
				default:
					t.Fatal("Not implemented")
				case *int:
					changed, concrete = setup(add, *v)
				case *int8:
					changed, concrete = setup(add, *v)
				case *int16:
					changed, concrete = setup(add, *v)
				case *int32:
					changed, concrete = setup(add, *v)
				case *int64:
					changed, concrete = setup(add, *v)
				case *uint:
					changed, concrete = setup(add, *v)
				case *uint8:
					changed, concrete = setup(add, *v)
				case *uint16:
					changed, concrete = setup(add, *v)
				case *uint32:
					changed, concrete = setup(add, *v)
				case *uint64:
					changed, concrete = setup(add, *v)
				case *uintptr:
					changed, concrete = setup(add, *v)
				case *float32:
					changed, concrete = setup(add, *v)
				case *float64:
					changed, concrete = setup(add, *v)
				case *complex64:
					changed, concrete = setup(add, *v)
				case *complex128:
					changed, concrete = setup(add, *v)
				case *string:
					changed, concrete = setup(add, *v)
				case *bool:
					changed, concrete = setup(func(b bool) bool { return !b}, *v)
				}
				assert.NotNil(t, got)
				t.Logf("%p %p", tc.data, got)
				assert.NotEqual(t, tc.data, got) // check pointers
				t.Logf("%v %v", changed, concrete)
				assert.NotEqual(t, changed, concrete) // check value
			default:
				got := deep.Copy(tc.data)
				t.Logf("%T %v", got, got)
				assert.Equal(t, tc.data, got)
			}
		})
	}
}

func setup[T any](mutate func(T) T, t T) (changed T, concrete T) { return mutate(t), t }

func toPtr[T any](t T) *T { return &t }

func add[T ordered](t T) T { return t + t }

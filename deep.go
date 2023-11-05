// Package deep deals with copying values down to the pointers. Many assignments
// will only do a "shallow" copy of two types. A shallow copy does not create
// new references to underlying pointers. This means that with a separate copied
// variable, you can still update values of the previously assigned variable.
// This is known, expected, and relied upon in many situations.
//
// This package is for when you want a completely separate copy of a type that
// has no attachment to the previous type. This is expensive as every value is
// allocated and copied into the new type. Use this, when necessary, not all the
// time.
//
// An example of a shallow copy:
//	a := [1, 2, 3, 4]
//	b := a
//	b[0] = 9
//	fmt.Println(a) // Output: [9, 2, 3, 4]
//
// An example of a deep copy:
//	a := [1, 2, 3, 4]
//	b := deep.Copy(a)
//	b[0] = 9
//	fmt.Println(a) // Output: [1, 2, 3, 4]
package deep

import (
	"fmt"
	"reflect"
)

// Copy creates a deep copy of original. A deep copy is defined as having no
// values inside of cpy that reference any value inside of original.
//
// In other words, all concrete values of original are the same as cpy, but any
// values that hold a reference (e.g. a pointer) are moved to a new reference
// with the same underlying concrete value.
func Copy[T any](original T) (cpy T) {
	val := reflect.ValueOf(original)
	var c T
	switch val.Kind() {
	case reflect.Pointer, reflect.UnsafePointer:
		c = reflect.New(reflect.TypeOf(original)).Interface().(T)
	case reflect.String, reflect.Bool, reflect.Uintptr,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128:
		c = val.Interface().(T)
	default:
		panic("Copy: not implemented!!!!" + fmt.Sprintf(" %T", original))
	}
	return c
}

// Copyright © 2013 Théo Crevon
//
// See the file LICENSE for copying permission.

// Package reflections provides high level abstractions over the Go standard [reflect] library.
//
// In practice, the `reflect` library's API proves somewhat low-level and un-intuitive.
// Using it can turn out pretty complex, daunting, and scary, when doing simple
// things like accessing a structure field value, a field tag, etc.
//
// The `reflections` package aims to make developers' life easier when it comes to introspect
// struct values at runtime. Its API takes inspiration in the python language's `getattr,` `setattr,` and `hasattr` set
// of methods and provides simplified access to structure fields and tags.
//
// [reflect]: http://golang.org/pkg/reflect/
package reflections

import (
	"errors"
	"reflect"
)

// ErrUnsupportedType indicates that the provided type doesn't support the requested reflection operation.
var ErrUnsupportedType = errors.New("unsupported type")

// ErrUnexportedField indicates that an operation failed as a result of
// applying to a non-exported struct field.
var ErrUnexportedField = errors.New("unexported field")

// GetField returns the value of the provided obj field.
// The `obj` can either be a structure or pointer to structure.
func GetField(obj interface{}, name string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFieldKind returns the kind of the provided obj field.
// The `obj` can either be a structure or pointer to structure.
func GetFieldKind(obj interface{}, name string) (reflect.Kind, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Kind), nil
}

// GetFieldType returns the kind of the provided obj field.
// The `obj` can either be a structure or pointer to structure.
func GetFieldType(obj interface{}, name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetFieldTag returns the provided obj field tag value.
// The `obj` parameter can either be a structure or pointer to structure.
func GetFieldTag(obj interface{}, fieldName, tagKey string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetFieldNameByTagValue looks up a field with a matching `{tagKey}:"{tagValue}"` tag in the provided `obj` item.
// The `obj` parameter must be a `struct`, or a `pointer` to one. If the `obj` parameter doesn't have a field tagged
// with the `tagKey`, and the matching `tagValue`, this function returns an error.
func GetFieldNameByTagValue(obj interface{}, tagKey, tagValue string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SetField sets the provided obj field with provided value.
//
// The `obj` parameter must be a pointer to a struct, otherwise it soundly fails.
// The provided `value` type should match with the struct field being set.
func SetField(obj interface{}, name string, value interface{}) error {
	_ = "STUB: not implemented"
	// Fetch the field reflect.Value
	return nil
}

// HasField checks if the provided `obj` struct has field named `name`.
// The `obj` can either be a structure or pointer to structure.
func HasField(obj interface{}, name string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Fields returns the struct fields names list.
// The `obj` parameter can either be a structure or pointer to structure.
func Fields(obj interface{}) ([]string, error) {
	_ = "STUB: not implemented"
	return nil,

		// FieldsDeep returns "flattened" fields.
		//
		// Note that FieldsDeep treats fields from anonymous inner structs as normal fields.
		nil
}

func FieldsDeep(obj interface{}) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func fields(obj interface{}, deep bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Items returns the field:value struct pairs as a map.
// The `obj` parameter can either be a structure or pointer to structure.
func Items(obj interface{}) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// ItemsDeep returns "flattened" items.
		// Note that ItemsDeep will treat fields from anonymous inner structs as normal fields.
		nil
}

func ItemsDeep(obj interface{}) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func items(obj interface{}, deep bool) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Tags lists the struct tag fields.
// The `obj` can whether be a structure or pointer to structure.
func Tags(obj interface{}, key string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// TagsDeep returns "flattened" tags.
	// Note that TagsDeep treats fields from anonymous
	// inner structs as normal fields.
}

func TagsDeep(obj interface{}, key string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func tags(obj interface{}, key string, deep bool) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func reflectValue(obj interface{}) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func isExportableField(field reflect.StructField) bool {
	_ = "STUB: not implemented"
	// PkgPath is empty for exported fields.
	return false
}

func isSupportedType(obj interface{}, types []reflect.Kind) bool {
	_ = "STUB: not implemented"
	return false
}

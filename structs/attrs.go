package structs

import (
	"fmt"
	"reflect"
)

// SetAttr set value to specified field.
func SetAttr(o interface{}, name string, value interface{}) error {
	r := reflect.ValueOf(o)
	if r.Kind() != reflect.Ptr || r.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("setattr: object must be a pointer to struct.")
	}
	f := r.Elem().FieldByName(name)
	if !f.IsValid() {
		return fmt.Errorf("setattr: no such field %s", name)
	}
	var v reflect.Value
	if value != nil {
		v = reflect.ValueOf(value)
	} else {
		v = reflect.Zero(f.Type())
	}
	if f.Kind() == reflect.Ptr && v.Kind() != reflect.Ptr {
		vf := reflect.New(reflect.TypeOf(value))
		vf.Elem().Set(v)
		v = vf
	} else if f.Kind() != reflect.Ptr && v.Kind() == reflect.Ptr {
		v = reflect.Indirect(v)
	}
	if f.Type() != v.Type() {
		return fmt.Errorf("setattr: field type mismatched. expected %s but actual %s.", f.Type(), v.Type())
	}
	f.Set(v)
	return nil
}

// GetAttr read value from specified field.
func GetAttr(o interface{}, name string, dst interface{}) error {
	r := reflect.ValueOf(o)
	if r.Kind() != reflect.Ptr || r.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("setattr: object must be a pointer to struct.")
	}
	f := r.Elem().FieldByName(name)
	if !f.IsValid() {
		return fmt.Errorf("setattr: no such field %s", name)
	}
	v := reflect.Indirect(reflect.ValueOf(dst))
	if !v.CanSet() {
		return fmt.Errorf("setattr: destination unassignable.")
	}
	if f.Kind() == reflect.Ptr {
		f = reflect.Indirect(f)
	}
	if f.Type() != v.Type() {
		return fmt.Errorf("setattr: field type mismatched. expected %s but actual %s.", f.Type(), v.Type())
	}
	v.Set(f)
	return nil
}

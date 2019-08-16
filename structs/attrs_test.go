package structs

import (
	"testing"
	"time"
)

//
func TestSetAttrInt(t *testing.T) {
	in := 100
	t.Run("value to value", func(t *testing.T) {
		o := &struct {
			F int
		}{}
		if err := SetAttr(o, "F", in); err != nil {
			t.Error(err)
		}
		if o.F != in {
			t.Errorf("expected %v but actual %v", in, o.F)
		}
	})
	t.Run("value to pointer", func(t *testing.T) {
		o := &struct {
			F *int
		}{}
		if err := SetAttr(o, "F", in); err != nil {
			t.Error(err)
		}
		if *o.F != in {
			t.Errorf("expected %v but actual %v", in, *o.F)
		}
	})
	t.Run("pointer to value", func(t *testing.T) {
		o := &struct {
			F int
		}{}
		if err := SetAttr(o, "F", &in); err != nil {
			t.Error(err)
		}
		if o.F != in {
			t.Errorf("expected %v but actual %v", in, o.F)
		}
	})
	t.Run("pointer to pointer", func(t *testing.T) {
		o := &struct {
			F *int
		}{}
		if err := SetAttr(o, "F", &in); err != nil {
			t.Error(err)
		}
		if *o.F != in {
			t.Errorf("expected %v but actual %v", in, *o.F)
		}
	})
}

//
func TestSetAttrFloat(t *testing.T) {
	in := 12.3
	t.Run("value to value", func(t *testing.T) {
		o := &struct {
			F float64
		}{}
		if err := SetAttr(o, "F", in); err != nil {
			t.Error(err)
		}
		if o.F != in {
			t.Errorf("expected %v but actual %v", in, o.F)
		}
	})
	t.Run("value to pointer", func(t *testing.T) {
		o := &struct {
			F *float64
		}{}
		if err := SetAttr(o, "F", in); err != nil {
			t.Error(err)
		}
		if *o.F != in {
			t.Errorf("expected %v but actual %v", in, *o.F)
		}
	})
	t.Run("pointer to value", func(t *testing.T) {
		o := &struct {
			F float64
		}{}
		if err := SetAttr(o, "F", &in); err != nil {
			t.Error(err)
		}
		if o.F != in {
			t.Errorf("expected %v but actual %v", in, o.F)
		}
	})
	t.Run("pointer to pointer", func(t *testing.T) {
		o := &struct {
			F *float64
		}{}
		if err := SetAttr(o, "F", &in); err != nil {
			t.Error(err)
		}
		if *o.F != in {
			t.Errorf("expected %v but actual %v", in, *o.F)
		}
	})
}

//
func TestSetAttrString(t *testing.T) {
	in := "qwerty"
	t.Run("value to value", func(t *testing.T) {
		o := &struct {
			F string
		}{}
		if err := SetAttr(o, "F", in); err != nil {
			t.Error(err)
		}
		if o.F != in {
			t.Errorf("expected %v but actual %v", in, o.F)
		}
	})
	t.Run("value to pointer", func(t *testing.T) {
		o := &struct {
			F *string
		}{}
		if err := SetAttr(o, "F", in); err != nil {
			t.Error(err)
		}
		if *o.F != in {
			t.Errorf("expected %v but actual %v", in, *o.F)
		}
	})
	t.Run("pointer to value", func(t *testing.T) {
		o := &struct {
			F string
		}{}
		if err := SetAttr(o, "F", &in); err != nil {
			t.Error(err)
		}
		if o.F != in {
			t.Errorf("expected %v but actual %v", in, o.F)
		}
	})
	t.Run("pointer to pointer", func(t *testing.T) {
		o := &struct {
			F *string
		}{}
		if err := SetAttr(o, "F", &in); err != nil {
			t.Error(err)
		}
		if *o.F != in {
			t.Errorf("expected %v but actual %v", in, *o.F)
		}
	})
}

//
func TestSetAttrTime(t *testing.T) {
	in := time.Now()
	t.Run("value to value", func(t *testing.T) {
		o := &struct {
			F time.Time
		}{}
		if err := SetAttr(o, "F", in); err != nil {
			t.Error(err)
		}
		if o.F != in {
			t.Errorf("expected %v but actual %v", in, o.F)
		}
	})
	t.Run("value to pointer", func(t *testing.T) {
		o := &struct {
			F *time.Time
		}{}
		if err := SetAttr(o, "F", in); err != nil {
			t.Error(err)
		}
		if *o.F != in {
			t.Errorf("expected %v but actual %v", in, *o.F)
		}
	})
	t.Run("pointer to value", func(t *testing.T) {
		o := &struct {
			F time.Time
		}{}
		if err := SetAttr(o, "F", &in); err != nil {
			t.Error(err)
		}
		if o.F != in {
			t.Errorf("expected %v but actual %v", in, o.F)
		}
	})
	t.Run("pointer to pointer", func(t *testing.T) {
		o := &struct {
			F *time.Time
		}{}
		if err := SetAttr(o, "F", &in); err != nil {
			t.Error(err)
		}
		if *o.F != in {
			t.Errorf("expected %v but actual %v", in, *o.F)
		}
	})
}

//
func TestAttrNil(t *testing.T) {
	t.Run("value to value", func(t *testing.T) {
		o := &struct {
			F string
		}{}
		if err := SetAttr(o, "F", nil); err != nil {
			t.Error(err)
		}
		if o.F != "" {
			t.Errorf("expected %v but actual %v", "", o.F)
		}
	})
	t.Run("value to pointer", func(t *testing.T) {
		o := &struct {
			F *string
		}{}
		if err := SetAttr(o, "F", nil); err != nil {
			t.Error(err)
		}
		if o.F != nil {
			t.Errorf("expected %v but actual %v", nil, o.F)
		}
	})
}

//
func TestGetAttrInt(t *testing.T) {
	t.Run("from value", func(t *testing.T) {
		var out int
		in := 100
		o := &struct {
			F int
		}{
			F: in,
		}
		if err := GetAttr(o, "F", &out); err != nil {
			t.Error(err)
		}
		if out != in {
			t.Errorf("expected %v but actual %v", in, out)
		}
	})
	t.Run("from pointer", func(t *testing.T) {
		var out int
		in := 100
		o := &struct {
			F *int
		}{
			F: &in,
		}
		if err := GetAttr(o, "F", &out); err != nil {
			t.Error(err)
		}
		if out != in {
			t.Errorf("expected %v but actual %v", in, out)
		}
	})
}

//
func TestGetAttrFloat(t *testing.T) {
	t.Run("from value", func(t *testing.T) {
		var out float64
		in := 12.3
		o := &struct {
			F float64
		}{
			F: in,
		}
		if err := GetAttr(o, "F", &out); err != nil {
			t.Error(err)
		}
		if out != in {
			t.Errorf("expected %v but actual %v", in, out)
		}
	})
	t.Run("from pointer", func(t *testing.T) {
		var out float64
		in := 12.3
		o := &struct {
			F *float64
		}{
			F: &in,
		}
		if err := GetAttr(o, "F", &out); err != nil {
			t.Error(err)
		}
		if out != in {
			t.Errorf("expected %v but actual %v", in, out)
		}
	})
}

//
func TestGetAttrTime(t *testing.T) {
	t.Run("from value", func(t *testing.T) {
		var out time.Time
		in := time.Now()
		o := &struct {
			F time.Time
		}{
			F: in,
		}
		if err := GetAttr(o, "F", &out); err != nil {
			t.Error(err)
		}
		if out != in {
			t.Errorf("expected %v but actual %v", in, out)
		}
	})
	t.Run("from pointer", func(t *testing.T) {
		var out time.Time
		in := time.Now()
		o := &struct {
			F *time.Time
		}{
			F: &in,
		}
		if err := GetAttr(o, "F", &out); err != nil {
			t.Error(err)
		}
		if out != in {
			t.Errorf("expected %v but actual %v", in, out)
		}
	})
}

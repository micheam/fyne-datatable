package datatable

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type myType struct {
	F1 string `column:"field_1" cc:"xxxxx"`
	F2 string `column:"field_2"`
	F3 string
	F4 string `column:"-"`
	F5 string `column:"field_5"`
}

func TestTagValues_SetTagkey(t *testing.T) {
	btgkey := tagkey
	SetTagkey("cc")
	t.Cleanup(func() { SetTagkey(btgkey) })
	want := []string{"xxxxx"}
	got := tagValues(myType{})
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("mismatch (-want, +got):%s\n", diff)
	}
}

func TestTagValues(t *testing.T) {
	want := []string{"field_1", "field_2", "field_5"}
	got := tagValues(myType{})
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("mismatch (-want, +got):%s\n", diff)
	}
}

func TestGetFieldValue(t *testing.T) {
	s := struct {
		F1 string `column:"-"`
		F2 string `column:"field_2"`
		F3 int    `column:"field_3"`
	}{
		F1: "xxxxx",
		F2: "yyyyy",
		F3: 11111,
	}
	t.Run("string field", func(t *testing.T) {
		want := "yyyyy"
		got := getFieldValue(s, 0)
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("mismatch (-want, +got):%s\n", diff)
		}
	})
	t.Run("int field", func(t *testing.T) {
		want := 11111
		got := getFieldValue(s, 1)
		if _, ok := got.(int); !ok {
			t.Errorf("want `int` but %T", got)
			t.FailNow()
		}
		if diff := cmp.Diff(got, want); diff != "" {
			t.Errorf("mismatch (-want, +got):%s\n", diff)
		}
	})
}

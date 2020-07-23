package testhelpers

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, exp, got interface{}, msg string) bool {
	if !reflect.DeepEqual(exp, got) {
		t.Errorf("expected %v, got: %v\n%s", exp, got, msg)
		return false
	}
	return true
}

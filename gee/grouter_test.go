package gee

import (
	"reflect"
	"testing"
)

func Test_newRouter(t *testing.T) {
	tests := []struct {
		name string
		want *router
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

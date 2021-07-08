package day3

import (
	"reflect"
	"testing"
)

func Test_combineToRouteKey(t *testing.T) {
	type args struct {
		methodType string
		pattern    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combineToRouteKey(tt.args.methodType, tt.args.pattern); got != tt.want {
				t.Errorf("combineToRouteKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newRouter(t *testing.T) {
	tests := []struct {
		name string
		want *router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parsePattern(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_router_addRoute(t *testing.T) {
	type fields struct {
		roots    map[string]*Node
		handlers map[string]HandlerFunc
	}
	type args struct {
		methodType string
		pattern    string
		handler    HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				roots:    tt.fields.roots,
				handlers: tt.fields.handlers,
			}
		})
	}
}

func Test_router_getRoute(t *testing.T) {
	type fields struct {
		roots    map[string]*Node
		handlers map[string]HandlerFunc
	}
	type args struct {
		methodType string
		pattern    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
		want1  map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				roots:    tt.fields.roots,
				handlers: tt.fields.handlers,
			}
			got, got1 := r.getRoute(tt.args.methodType, tt.args.pattern)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRoute() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getRoute() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_router_handle(t *testing.T) {
	type fields struct {
		roots    map[string]*Node
		handlers map[string]HandlerFunc
	}
	type args struct {
		c *Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				roots:    tt.fields.roots,
				handlers: tt.fields.handlers,
			}
		})
	}
}

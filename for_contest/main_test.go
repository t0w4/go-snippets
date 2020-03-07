package main

import (
	"reflect"
	"testing"
)

func Test_sum(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "normal 1", args: args{a: []int{1, 2, 3, 4, 5}}, want: 15},
		{name: "normal 2", args: args{a: []int{-1, 2, -3, 4, -5}}, want: -3},
		{name: "normal 3", args: args{a: []int{8, 2, 6, 4, 5, 2, 5, 1}}, want: 33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.a); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unique(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "normal 1", args: args{a: []string{"a", "a", "A", "B", "C"}}, want: []string{"a", "A", "B", "C"}},
		{name: "normal 2", args: args{a: []string{"C", "B", "A", "a", "a"}}, want: []string{"C", "B", "A", "a"}},
		{name: "normal 3", args: args{a: []string{"a", "a", "b", "b", "c"}}, want: []string{"a", "b", "c"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unique(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

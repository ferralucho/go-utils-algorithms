package main

import (
	"testing"
)

func Test_compute(t *testing.T) {

	type args struct {
		fn func(float64, float64) float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compute(tt.args.fn); got != tt.want {
				t.Errorf("compute() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func TestIsBalancedPairs(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"manyBalanced", args{s: "([])(){}(())()()"}, true},
		{"manyUnbalanced", args{s: "{[[[[({(}))]]]]}"}, false},
		{"manyWordsBalanced", args{s: "([])(dsd){}(())()()"}, true},
		{"manyWordsUnbalanced", args{s: "{[[[[({(asdqwe}))]]]ee]}"}, false},
		{"oneRuneUnbalanced", args{s: "["}, false},
		{"empty", args{s: ""}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBalancedPairs(tt.args.s); got != tt.want {
				t.Errorf("IsBalancedPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

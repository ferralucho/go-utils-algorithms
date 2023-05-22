package main

import (
	"reflect"
	"testing"
)

func TestPostOrderTraversal(t *testing.T) {
	root := &Node{
		data: 1,
		left: &Node{
			data: 2,
			left: &Node{
				data: 4,
			},
			right: &Node{
				data: 5,
			},
		},
		right: &Node{
			data: 3,
			left: &Node{
				data: 6,
			},
			right: &Node{
				data: 7,
			},
		},
	}

	result := root.PostOrderTraversal()

	expected := []int{4, 5, 2, 6, 7, 3, 1}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

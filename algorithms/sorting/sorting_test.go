package sorting

import (
	"reflect"
	"testing"
)

func Test_InsertionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name           string
		args           args
		wantSorted_arr []int
	}{
		{"#0 Empty array", args{[]int{}}, []int{}},
		{"#1", args{[]int{1, 0}}, []int{0, 1}},
		{"#2", args{[]int{1, 0, 3}}, []int{0, 1, 3}},
		{"#3", args{[]int{1, 0, 3, 2}}, []int{0, 1, 2, 3}},
		{"#4", args{[]int{3, 2, 1, 4, 7, 8, 6, 9, 5, 0}}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSorted_arr := InsertionSort(tt.args.arr); !reflect.DeepEqual(gotSorted_arr, tt.wantSorted_arr) {
				t.Errorf("InsertionSort() = %v, want %v", gotSorted_arr, tt.wantSorted_arr)
			}
		})
	}
}

func Test_Revert(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"#0 Empty array", args{[]int{}}, []int{}},
		{"#1", args{[]int{1, 0}}, []int{0, 1}},
		{"#2", args{[]int{1, 0, 3}}, []int{3, 0, 1}},
		{"#3", args{[]int{1, 0, 3, 2}}, []int{2, 3, 0, 1}},
		{"#4", args{[]int{3, 2, 1, 4, 7, 8, 6, 9, 5, 0}}, []int{0, 5, 9, 6, 8, 7, 4, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Revert(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Revert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_InsertionSort_BinaryTree(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name           string
		args           args
		wantSorted_arr []int
	}{
		{"#0 Empty array", args{[]int{}}, []int{}},
		{"#1", args{[]int{1, 0}}, []int{0, 1}},
		{"#2", args{[]int{1, 0, 3}}, []int{0, 1, 3}},
		{"#3", args{[]int{1, 0, 3, 2}}, []int{0, 1, 2, 3}},
		{"#4", args{[]int{3, 2, 1, 4, 7, 8, 6, 9, 5, 0}}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSorted_arr := InsertionSort_BinaryTree(tt.args.arr); !reflect.DeepEqual(gotSorted_arr, tt.wantSorted_arr) {
				t.Errorf("InsertionSort_BinaryTree() = %v, want %v", gotSorted_arr, tt.wantSorted_arr)
			}
		})
	}
}

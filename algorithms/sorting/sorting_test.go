package sorting

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
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

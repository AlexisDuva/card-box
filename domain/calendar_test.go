package domain

import (
	"reflect"
	"testing"
)

func TestAssessment(t *testing.T) {
	tests := []struct {
		name string
		nbc  int
		age  int
		want []int
	}{
		{
			name: "no cells",
			nbc:  0, age: 1,
			want: []int{},
		},
		{
			name: "cell 0 is tested every day (cycle 1)",
			nbc:  1, age: 3,
			want: []int{0},
		},
		{
			name: "day 1: only cell 0 (1 is not divisible by 2, 4...)",
			nbc:  4, age: 1,
			want: []int{0},
		},
		{
			name: "day 2: cells 0 and 1 (cycle 2)",
			nbc:  4, age: 2,
			want: []int{0, 1},
		},
		{
			name: "day 3: only cell 0",
			nbc:  4, age: 3,
			want: []int{0},
		},
		{
			name: "day 4: cells 0, 1 and 2 (cycle 4)",
			nbc:  4, age: 4,
			want: []int{0, 1, 2},
		},
		{
			name: "day 8: all 4 cells (cycle 8)",
			nbc:  4, age: 8,
			want: []int{0, 1, 2, 3},
		},
		{
			name: "day 6: cells 0 and 1 (6 divisible by 1 and 2, not 4)",
			nbc:  4, age: 6,
			want: []int{0, 1},
		},
		{
			name: "nbc limits which cells are assessed",
			nbc:  2, age: 8,
			want: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Assessment(tt.nbc, tt.age)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Assessment(%d, %d) = %v, want %v", tt.nbc, tt.age, got, tt.want)
			}
		})
	}
}

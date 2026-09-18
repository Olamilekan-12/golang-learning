package calc

import "testing"

func TestAdd(t *testing.T) {
	test := []struct {
		name string
		a, b int
		want int
	}{
		{"postive numbers", 2, 3, 5},
		{"with zero", 0, 7, 7},
		{"negative numbers", -4, -6, -10},
		{"mixed signs", -5, 5, 0},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %xd) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

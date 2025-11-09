package domain

import "testing"

func TestInBorders(t *testing.T) {
	tests := []struct {
		name     string
		h, w     int
		row, col int
		expected bool
	}{
		// inner 1x1 -> scaled 3x3
		{"1x1_top_left", 1, 1, 0, 0, true},
		{"1x1_bottom_right", 1, 1, 2, 2, true},
		{"1x1_beyond_row", 1, 1, 3, 0, false},
		{"1x1_negative_col", 1, 1, 0, -1, false},

		// inner 2x2 -> scaled 5x5
		{"2x2_center", 2, 2, 2, 2, true},
		{"2x2_edge_bottom", 2, 2, 4, 2, true},
		{"2x2_row_eq_scaledH", 2, 2, 5, 1, false},
		{"2x2_col_eq_scaledW", 2, 2, 1, 5, false},

		// inner 3x4 -> scaled 7x9
		{"3x4_inside", 3, 4, 6, 8, true},
		{"3x4_beyond_br", 3, 4, 7, 9, false},

		// edge conditions
		{"0x0_scaled_1x1_origin", 0, 0, 0, 0, true},
		{"0x0_scaled_1x1_out", 0, 0, 1, 0, false},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			m := NewMaze(tc.h, tc.w)

			got := m.InBorders(tc.row, tc.col)
			if got != tc.expected {
				t.Fatalf("inBorders(%d,%d) on scaled %dx%d (inner %dx%d) = %v; want %v",
					tc.row, tc.col, 2*tc.h+1, 2*tc.w+1, tc.h, tc.w, got, tc.expected)
			}
		})
	}
}

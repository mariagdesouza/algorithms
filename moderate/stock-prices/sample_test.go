package stockprices

import "testing"

func TestFindBestPrice(t *testing.T) {

	testTable := []struct {
		data     []float64
		expected float64
	}{
		{[]float64{110.55, 112.50, 130.25, 120.25, 145.75, 110.00, 90.90}, 35.2},
		{[]float64{150.55, 120.50, 115.25, 111.25, 90.75, 30.00, 20.90}, -4},
		{[]float64{110.75, 112.50, 130.25, 120.25, 145.75, 110.00, 90.90}, 35},
		{[]float64{5, 4, 3, 4, 5}, 2},
	}
	for _, test := range testTable {
		result := FindBestPrice(test.data)
		if result != test.expected {
			t.Errorf("FindBestPrice failed. returned = %v; want: %v", result, test.expected)
		}
	}
}

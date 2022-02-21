package calc

import (
	"fmt"
	"testing"
)

//arr = []int32{1,2,3,4,5,6}
func TestNumbersInStringCalculator(t *testing.T) {

	var nums = []struct {
		inp  []string
		outp []float64 //verified
	}{
		{[]string{"500-450-65-85", "12+55+89+100", "1500/3/2/2", "3*67*2*2"},
			[]float64{-100, 256, 125, 804}},
		{[]string{"30-2-3-15", "100+10+11+15", "100/1/2/2", "3*2*2*1"},
			[]float64{10, 136, 25, 12}},
		{[]string{"600-150-85-85", "120+55+45+100", "1200/4/3/2", "4*12*7*5"},
			[]float64{280, 320, 50, 1680}},
	}

	for _, inputs := range nums {
		floatResult := NumbersInStringCalculator(inputs.inp...)
		for i := 0; i < len(inputs.outp); i++ {
			if floatResult[i] != inputs.outp[i] {
				t.Errorf("Expected output %v received %v", inputs.outp, floatResult)
				fmt.Println(floatResult[i])
			}
		}

	}
}

func TestAddition(t *testing.T) {

	var nums = []struct {
		inp  []float64
		outp float64 //verified output
	}{
		{[]float64{4, 5, 6, 2}, 17},
		{[]float64{2, 7, 2, 3}, 14},
	}

	for _, inputs := range nums {
		floatResult := Addition(inputs.inp)
		if floatResult != inputs.outp {
			t.Errorf("Addition : Expected %v : output %v", inputs.outp, floatResult)
		}
	}

}

func TestSubstract(t *testing.T) {
	var nums = []struct {
		inp  []float64
		outp float64
	}{
		{[]float64{100 - 20 - 30 - 5}, 45},
		{[]float64{200 - 90 - 10 - 20}, 80},
	}

	for _, inputs := range nums {
		floatResult := Substract(inputs.inp)
		if floatResult != inputs.outp {
			t.Errorf("Substract : Expected %v : output %v", inputs.outp, floatResult)
		}

	}
}

func TestDivision(t *testing.T) {
	var nums = []struct {
		inp  []float64
		outp float64
	}{
		{[]float64{1000, 5, 2, 2}, 50},
		{[]float64{2000, 5, 2, 2}, 100},
	}
	for _, inputs := range nums {
		floatResult := Division(inputs.inp)
		if floatResult != inputs.outp {
			t.Errorf("Division: Expected Output %v : output %v", inputs.outp, floatResult)
		}
	}
}

func TestMultiply(t *testing.T) {
	var nums = []struct {
		inp  []float64
		outp float64
	}{
		{[]float64{3, 2, 5, 6}, 180},
		{[]float64{4, 2, 6, 7}, 336},
	}

	for _, inputs := range nums {
		floatResult := Multiply(inputs.inp)
		if floatResult != inputs.outp {
			t.Errorf("Multiply: Expected Output %v: output %v", inputs.outp, floatResult)
		}
	}
}

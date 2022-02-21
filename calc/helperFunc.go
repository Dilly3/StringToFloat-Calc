package calc

import (
	"fmt"
	"strconv"
)

//type gbenga = []float64
type flt64 = []float64 // create type of  float64 slice

// ***************************** MULTIPLY FUNCTION **************************

func Multiply(arr flt64) float64 {
	// [2,3,4,5]
	var multiplyValue float64 = 1
	for _, val := range arr {

		multiplyValue = multiplyValue * val
	}
	multiplyValue, _ = strconv.ParseFloat(fmt.Sprintf("%0.6f", multiplyValue), 64)
	return multiplyValue
}

//******************************ADDITION FUNCTION *************************

func Addition(arr flt64) float64 {

	var additionValue float64 = 0
	for _, val := range arr {

		additionValue = additionValue + val
	}
	return additionValue
}

// *********************************** DIVISION FUNCTION ***************************

func Division(arr flt64) float64 {

	var divisionValue float64 = arr[0]
	for i := 1; i < len(arr); i++ {

		divisionValue = divisionValue / arr[i]
	}
	divisionValue, _ = strconv.ParseFloat(fmt.Sprintf("%f", divisionValue), 64)
	return divisionValue
}

//************************************** SUBTRACTION FUNCTION ***************************

func Substract(arr flt64) float64 {

	var subtractionValue = arr[0]
	for i := 1; i < len(arr); i++ {

		subtractionValue = subtractionValue - arr[i]
	}
	return subtractionValue
}

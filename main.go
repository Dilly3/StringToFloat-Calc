package main

import (
	"fmt"
	ohdy "week2/dilly3/calc"
)

func init() {

}
func main() {
	//custom type

	type floatArray = []float64
	arr := ohdy.NumbersInStringCalculator("500-450-65-85", "12+55+89+100", "1200-340-52-257")
	fmt.Println(arr)
	//floater := []float64{34, 56, 78, 98}
	fmt.Println(ohdy.Multiply(floatArray{34, 5, 78, 6}))
	fmt.Println(ohdy.Addition(floatArray{34, 56, 78, 56}))
	fmt.Println(ohdy.Division(floatArray{3404000, 56, 78, 56}))
	fmt.Println(ohdy.Substract(floatArray{304, 56, 78, 56}))

	//floatArr := []float64{}
	//floatArr = append(floatArr, 45)
	//fmt.Println(floatArr)
}

package calc

import (
	"strconv"
	"strings"
)

//  calculate( “2*3*4*5” , “25/5/2”, “2+3+4+5.9+6+7.8”,“20.54-7.65-2.897”)

//should return [120, 2.5, 28.7, 9.]

func NumbersInStringCalculator(args ...string) []float64 {

	var resultArr []float64 = make([]float64, len(args))

	for i := 0; i < len(args); i++ {

		switch {
		case strings.Contains(args[i], "*"):

			// initial string = “2*3*4*5” ==>  [2,3,4,5]//float slice
			{
				stringArr := strings.Split(args[i], "*")
				var floatArr []float64
				for i := 0; i < len(stringArr); i++ {
					newI, _ := strconv.ParseFloat(stringArr[i], 64)
					floatArr = append(floatArr, newI)
				}

				resultArr[i] = Multiply(floatArr)
				break
			}

		case strings.Contains(args[i], "+"):
			{
				// 2+3+4+5.9+6+7.8
				stringArr := strings.Split(args[i], "+")
				floatArr := []float64{}
				for i := 0; i < len(stringArr); i++ {
					newI, _ := strconv.ParseFloat(stringArr[i], 64)
					floatArr = append(floatArr, newI)
				}
				resultArr[i] = Addition(floatArr)
				break
			}

		case strings.Contains(args[i], "/"):
			{
				stringArr := strings.Split(args[i], "/")
				floatArr := []float64{}
				for i := 0; i < len(stringArr); i++ {
					newI, _ := strconv.ParseFloat(stringArr[i], 64)
					floatArr = append(floatArr, newI)
				}
				resultArr[i] = Division(floatArr)
				break
			}
		case strings.Contains(args[i], "-"):
			{
				stringArr := strings.Split(args[i], "-")
				floatArr := []float64{}
				for i := 0; i < len(stringArr); i++ {
					newI, _ := strconv.ParseFloat(stringArr[i], 64)
					floatArr = append(floatArr, newI)
				}
				resultArr[i] = Substract(floatArr)
			}

		}
	}
	//for _, v := range arr {
	//	fmt.Println(v)
	//}

	return resultArr

}

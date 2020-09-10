package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	total := 0
	for _, val := range numbers {
		total += val
	}
	return total
}

//Recursion way
//func Sum(numbers []int) int  {
//	if len(numbers) == 0 {
//		return 0
//	}
//	return numbers[0] + Sum(numbers[1:])
//}

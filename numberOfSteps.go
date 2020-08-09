package number_of_steps

func numberOfSteps(num int) int {
	count := 0
	for num > 0 {
		if num%2 == 0 {
			num = num / 2
		} else {
			num -= 1
		}
		count++
	}
	return count
}

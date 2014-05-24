package divisor

func HigestPrime(num int) int {
	for num%2 == 0 {
		num /= 2
	}
	if num == 1 {
		return (2)
	}
	for i := 3; i <= num; i += 2 {
		for num%i == 0 {
			num /= i
		}
		if num == 1 {
			return (i)
		}
	}
	return 0
}

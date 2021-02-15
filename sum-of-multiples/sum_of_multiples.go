package summultiples

func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	set := make(map[int]bool)
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if d == 0 {
				continue
			}
			if i%d == 0 {
				if _, ok := set[i]; !ok {
					sum += i
					set[i] = true
				}
			}
		}
	}
	return sum
}

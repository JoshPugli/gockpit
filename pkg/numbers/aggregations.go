package numbers

import "fmt"

func Max[T Number](numbers ...T) (T, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no arguments provided")
	}

	maxVal := numbers[0]
	for _, val := range numbers[1:] {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal, nil
}

func Min[T Number](numbers ...T) (T, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no arguments provided")
	}

	minVal := numbers[0]
	for _, val := range numbers[1:] {
		if val < minVal {
			minVal = val
		}
	}

	return minVal, nil
}

func Sum[T Number](numbers ...T) (T, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no arguments provided")
	}

	var sum T
	for _, val := range numbers {
		sum += val
	}

	return sum, nil
}

func Average[T Number](numbers ...T) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no arguments provided")
	}

	sum, err := Sum(numbers...)
	if err != nil {
		return 0, err
	}

	average := float64(sum) / float64(len(numbers))
	return average, nil
}

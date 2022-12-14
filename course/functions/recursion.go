package main

// func factorial(n int) (int, error) {
// 	switch {
// 	case n < 0:
// 		return -1, fmt.Errorf("invalid number: %d", n)
// 	case n == 0:
// 		return 1, nil
// 	default:
// 		previous, _ := factorial(n - 1)
// 		return n * previous, nil
// 	}
// }

func factorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * factorial(n-1)
	}
}

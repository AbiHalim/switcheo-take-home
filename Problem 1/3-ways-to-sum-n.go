func sum_to_n_a(n int) int {
	// O(1) time complexity from arithmetic computation
	if n%2 == 0 {
		return (n + 1) * n / 2
	} else {
		return (n+1)*(n/2) + (n+1)/2
	}
}

func sum_to_n_b(n int) int {
	// O(n) time complexity from iterating from 1 to nx
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func sum_to_n_c(n int) int {
	// O(n) time complexity from recursing down to 1
	if n == 0 {
		return 0
	} else {
		return n + sum_to_n_c(n-1)
	}
}

// Edge case: the problem is ill-defined for negative integers
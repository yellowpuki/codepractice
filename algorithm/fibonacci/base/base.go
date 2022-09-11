package base

func Fibo(n int) int {
	if n <= 2 {
		return n
	}
	return Fibo(n-1) + Fibo(n-2)
}

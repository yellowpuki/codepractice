package advanced

var (
	cache map[int]int
)

func init() {
	cache = make(map[int]int)
}

func CachedFibo(n int) int {
	if n <= 2 {
		return n
	}

	if v, ok := cache[n]; ok {
		return v
	}

	cache[n] = CachedFibo(n-1) + CachedFibo(n-2)

	return cache[n]
}

func KMP(w []byte) []int {
	n, k := len(w)+1, 0
	p := make([]int, n)
	p[1] = 0
	for i := 2; i < n; i++ {
		for k > 0 && w[k] != w[i-1] {
			k = p[k]
		}
		if w[k] == w[i-1] {
			k++
		}
		p[i] = k
	}
	return p
}

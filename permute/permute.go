package permute

func Permutations[T any](ts []T) [][]T {
	var recurse func([]T, int)
	var perms [][]T

	recurse = func(ts []T, n int) {
		if n == 1 {
			tmp := make([]T, len(ts))
			copy(tmp, ts)
			perms = append(perms, tmp)
		} else {
			for i := 0; i < n; i++ {
				recurse(ts, n-1)

				if n%2 == 1 {
					tmp := ts[i]
					ts[i] = ts[n-1]
					ts[n-1] = tmp
				} else {
					tmp := ts[0]
					ts[0] = ts[n-1]
					ts[n-1] = tmp
				}
			}
		}
	}

	recurse(ts, len(ts))

	return perms
}

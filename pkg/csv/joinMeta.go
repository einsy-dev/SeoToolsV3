package csv

func joinMeta(dst map[string]int, src map[string]int) map[string]int {

	for k := range src {
		_, ok := dst[k]
		if !ok {
			dst[k] = len(dst)
			continue
		}
	}

	return dst
}

package csv

func join(dst *CSV, src *CSV) {
	dst.cols = joinMeta(dst.cols, src.cols)
	dst.rows = joinMeta(dst.rows, src.rows)

	maxRows := len(dst.rows)
	maxCols := len(dst.cols)

	dst.Data = append(dst.Data, make([][]string, maxRows-len(dst.Data))...)

	for i := range dst.Data {
		dst.Data[i] = append(dst.Data[i], make([]string, maxCols-len(dst.Data[i]))...)
	}

	for r := range src.rows {
		for c := range src.cols {
			dst.set(r, c, src.get(r, c))
		}
	}

}

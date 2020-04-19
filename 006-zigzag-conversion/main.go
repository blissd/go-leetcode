package main

type cell int

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	//width/height of a diagonal
	dlen := numRows - 2
	colStep := numRows + dlen

	ss := make([]uint8, len(s), len(s))
	for i := 0; i < len(ss); i++ {
		ss[i] = '.'
	}

	// i = position to write to in ss
	// diagStep = distance from column to diagonal value
	diagStep := colStep - 2
	i := 0
	for row := 0; row < numRows && i < len(ss); row = row + 1 {
		// does this row have diagonal cells?
		diag := row > 0 && row < numRows-1
		for j := row; j < len(ss) && i < len(ss); j += colStep {
			ss[i] = s[j]
			i += 1
			if diag && j+diagStep < len(s) {
				ss[i] = s[j+diagStep]
				i += 1
			}
		}
		if diag {
			diagStep -= 2
		}
	}

	return string(ss)
}

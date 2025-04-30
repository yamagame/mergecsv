package main

func mergeCSVData(data1, data2 [][]string, update bool) [][]string {
	if len(data1) == 0 || len(data2) == 0 {
		return data1
	}

	header1 := data1[0]
	header2 := data2[0]

	columnMap := make(map[int]int)
	for i, col1 := range header1 {
		for j, col2 := range header2 {
			if col1 == col2 {
				columnMap[i] = j
				break
			}
		}
	}

	for i := 1; i < len(data1); i++ {
		for j, colIndex2 := range columnMap {
			if i < len(data2) && colIndex2 < len(data2[i]) {
				if update {
					if data2[i][colIndex2] != "" {
						data1[i][j] = data2[i][colIndex2]
					}
				} else {
					if data1[i][j] == "" {
						data1[i][j] = data2[i][colIndex2]
					}
				}
			}
		}
	}

	return data1
}

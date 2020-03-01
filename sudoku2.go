package main

import "fmt"

var grid = [9][9]int{{0, 9, 6, 0, 4, 0, 0, 0, 1}, {1, 0, 0, 0, 6, 0, 0, 0, 4}, {5, 0, 4, 8, 1, 0, 3, 9, 0},
 					 {0, 0, 7, 9, 5, 0, 0, 4, 3}, {0, 3, 0, 0, 8, 0, 0, 0, 0}, {4, 0, 5, 0, 2, 3, 0, 1, 8},
 					 {0, 1, 0, 6, 3, 0, 0, 5, 9}, {0, 5, 9, 0, 7, 0, 8, 3, 0}, {0, 0, 3, 5, 9, 0, 0, 0, 7}}

func main() {


	fmt.Println(checkInput(grid))
	fmt.Println(grid)


}

func nextEmpty(i int, j int) (int, int) {
	for x := i; x < 9; x++ {
		for y := j; y < 9; y++ {
			if grid[x][y] == 0 {
				return x, y
			}
		}
	}
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if grid[x][y] == 0 {
				return x, y
			}
		}
	}
	return -1, -1
}

func isValid(i int, j int, e int) bool {
	rowCheck := true
	for x := 0; x < 9; x++ {
		if grid[i][x] == e {
			rowCheck = false
		}
	}
	if rowCheck {
		columnCheck := true
		for x := 0; x < 9; x++ {
			if grid[x][j] == e {
				columnCheck = false
			}
		}
		if columnCheck {
			secTopX, secTopY := 3*(i/3), 3*(j/3)
			for x := secTopX; x < secTopX+3; x++ {
				for y := secTopY; y < secTopY+3; y++ {
					if grid[x][y] == e {
						return false
					}
				}
			}
			return true
		}
	}
	return false
}

func solveSudoku(i int, j int) bool {
	i, j = nextEmpty(i, j)
	if i == -1 {
		return true
	}
	for e := 1; e < 10; e++ {
		if isValid(i, j, e) {
			grid[i][j] = e
			if solveSudoku(i, j) {
				return true
			}
			grid[i][j] = 0
		}
	}
	return false
}

func checkEachInput(grid [9][9]int, i int, j int, e int) bool {

	if (grid[i][j]!=0){
		counter := 0
		rowCheck := true
			for x := 0; x < 9; x++ {
				if grid[i][x] == e {
					counter++
					if counter>1{
						rowCheck = false
						counter = 0
						return false
						break
					}
				}
			}
			if rowCheck {
				columnCheck := true
				counter= 0
				for x := 0; x < 9; x++ {
					if grid[x][j] == e {
						counter++
						if counter>1{
						columnCheck = false
						return false
						counter = 0
						break
					}
					}
				}
				if columnCheck {
					secTopX, secTopY := 3*(i/3), 3*(j/3)
					counter=0
					for x := secTopX; x < secTopX+3; x++ {
						for y := secTopY; y < secTopY+3; y++ {
							if grid[x][y] == e {
								counter++
								if counter >1{
									counter=0
									return false
								}
							}
						}
					}
				}
			}
	}
	return true
}

func checkInput(grid [9][9]int) bool{
		for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if !(checkEachInput(grid, x,y, grid[x][y])){
				return false
			}
		}
	}
	return true
}
package main

func getGroups(grid [][]bool) [][][]int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid))
	}
	groups := [][][]int{}
	for column := 0; column < len(grid); column++ {
		for row := 0; row < len(grid[0]); row++ {
			group := getGroup(row, column, grid, visited, [][]int{})
			if len(group) > 0 {
				groups = append(groups, group)
			}
		}
	}
	return groups
}

func getGroup(row, column int, grid [][]bool, visited [][]bool, group [][]int) [][]int {
	if row < 0 || column < 0 || row >= len(grid[0]) || column >= len(grid) {
		return group
	}
	if !grid[column][row] || visited[column][row] {
		return group
	}
	visited[column][row] = true
	group = append(group, []int{column, row})
	group = getGroup(row-1, column-1, grid, visited, group)
	group = getGroup(row, column-1, grid, visited, group)
	group = getGroup(row+1, column-1, grid, visited, group)
	group = getGroup(row-1, column, grid, visited, group)
	group = getGroup(row+1, column, grid, visited, group)
	group = getGroup(row-1, column+1, grid, visited, group)
	group = getGroup(row, column+1, grid, visited, group)
	group = getGroup(row+1, column+1, grid, visited, group)
	return group
}

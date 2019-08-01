package gridgrouper

func getGroups(grid BoolGrid) [][]Pos {
	var groups [][]Pos
	if len(grid) == 0 {
		return groups
	}
	visited := NewBoolGrid(len(grid), len(grid[0]))
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			group := getGroup(row, col, grid, visited, []Pos{})
			if len(group) > 0 {
				groups = append(groups, group)
			}
		}
	}
	return groups
}

func getGroup(row, col int, grid, visited BoolGrid, group []Pos) []Pos {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[0]) {
		return group
	}
	if !grid[row][col] || visited[row][col] {
		return group
	}
	visited[row][col] = true
	group = append(group, Pos{row, col})
	group = getGroup(row-1, col-1, grid, visited, group)
	group = getGroup(row, col-1, grid, visited, group)
	group = getGroup(row+1, col-1, grid, visited, group)
	group = getGroup(row-1, col, grid, visited, group)
	group = getGroup(row+1, col, grid, visited, group)
	group = getGroup(row-1, col+1, grid, visited, group)
	group = getGroup(row, col+1, grid, visited, group)
	group = getGroup(row+1, col+1, grid, visited, group)
	return group
}

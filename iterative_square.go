package main

type position struct {
	Column int
	Row    int
}

type group = []position

func getGroupsBySquares(grid [][]bool) []group {
	// Group into 2x2 squares
	n, m := len(grid), len(grid[0])
	groups := [][]**group{}
	for i := 0; i < n; i += 2 {
		row := []**group{}
		for j := 0; j < m; j += 2 {
			g := &group{}
			if grid[i][j] {
				*g = append(*g, position{j, i})
			}
			if j+1 < m && grid[i][j+1] {
				*g = append(*g, position{j + 1, i})
			}
			if i+1 < n && grid[i+1][j] {
				*g = append(*g, position{j, i + 1})
			}
			if i+1 < n && j+2 < m && grid[i+1][j+1] {
				*g = append(*g, position{j + 1, i + 1})
			}
			row = append(row, &g)
		}
		groups = append(groups, row)
	}

	// Consolidate square groups
	for i := 0; i < len(groups); i++ {
		for j := 0; j < len(groups[0]); j++ {
			if len(**groups[i][j]) == 0 {
				continue
			}
			gi, gj := i*2, j*2
			if j+1 < len(groups[0]) {
				// Combine with group to the right
				if groups[i][j] != groups[i][j+1] &&
					(grid[gi][gj+1] || grid[gi+1][gj+1]) &&
					(grid[gi][gj+2] || grid[gi+1][gj+2]) {
					mergeGroups(groups, i, j, i, j+1)
				}
			}
			if i+1 < len(groups) {
				// Combine with group to the bottom-left
				if j > 0 && groups[i][j] != groups[i+1][j-1] &&
					grid[gi+1][gj] && grid[gi+2][gj-1] {
					mergeGroups(groups, i, j, i+1, j-1)
				}
				// Combine with group below
				if groups[i][j] != groups[i+1][j] &&
					(grid[gi+1][gj] || grid[gi+1][gj+1]) &&
					(grid[gi+2][gj] || grid[gi+2][gj+1]) {
					mergeGroups(groups, i, j, i+1, j)
				}
				// Combine with group to the bottom-right
				if j+1 < len(groups[0]) && groups[i][j] != groups[i+1][j+1] &&
					grid[gi+1][gj+1] && grid[gi+2][gj+2] {
					mergeGroups(groups, i, j, i+1, j+1)
				}
			}
		}
	}

	uniqueGroups := map[*group]bool{}
	finalGroups := []group{}
	for i := 0; i < len(groups); i++ {
		for j := 0; j < len(groups[0]); j++ {
			if !uniqueGroups[*groups[i][j]] && len(**groups[i][j]) > 0 {
				uniqueGroups[*groups[i][j]] = true
				finalGroups = append(finalGroups, **groups[i][j])
			}
		}
	}

	return finalGroups
}

func mergeGroups(groups [][]**group, i1, j1, i2, j2 int) {
	**groups[i1][j1] = append(**groups[i1][j1], **groups[i2][j2]...)
	*groups[i2][j2] = *groups[i1][j1]
	groups[i2][j2] = groups[i1][j1]
}

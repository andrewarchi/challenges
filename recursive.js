function getGroups(grid) {
  const groups = [];
  if (grid.length === 0) {
    return groups
  }
  const visited = [...Array(grid.length)].map(row => Array(grid[0].length).fill(false));
  for (let row = 0; row < grid.length; row++) {
    for (let col = 0; col < grid[0].length; col++) {
      const group = getGroup(row, col, grid, visited, []);
      if (group.length > 0) {
        groups.push(group);
      }
    }
  }
  return groups;
}

function getGroup(row, col, grid, visited, group) {
  if (row < 0 || col < 0 || row >= grid.length || col >= grid[0].length) {
    return group;
  }
  if (!grid[row][col] || visited[row][col]) {
    return group;
  }
  visited[row][col] = true;
  group.push([row, col]);
  getGroup(row-1, col-1, grid, visited, group);
  getGroup(row-1, col,   grid, visited, group);
  getGroup(row-1, col+1, grid, visited, group);
  getGroup(row,   col-1, grid, visited, group);
  getGroup(row,   col+1, grid, visited, group);
  getGroup(row+1, col-1, grid, visited, group);
  getGroup(row+1, col,   grid, visited, group);
  getGroup(row+1, col+1, grid, visited, group);
  return group;
}

console.log(getGroups([
  [true, false, true, true],
  [false, false, true, false],
  [true, true, false, false],
  [true, false, false, false]
]));

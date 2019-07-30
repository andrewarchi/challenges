function getGroups(grid) {
  const visited = [...Array(grid.length)].map(row => Array(grid[0].length).fill(false));
  const groups = [];
  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[0].length; j++) {
      const group = getGroup(i, j, grid, visited, []);
      if (group.length) {
        groups.push(group);
      }
    }
  }
  return groups;
}

function getGroup(i, j, grid, visited, group) {
  if (i < 0 || j < 0 || i >= grid.length || j >= grid[0].length) {
    return group;
  }
  if (!grid[i][j] || visited[i][j]) {
    return group;
  }
  visited[i][j] = true;
  group.push([i, j]);
  getGroup(i-1, j-1, grid, visited, group);
  getGroup(i-1, j,   grid, visited, group);
  getGroup(i-1, j+1, grid, visited, group);
  getGroup(i,   j-1, grid, visited, group);
  getGroup(i,   j+1, grid, visited, group);
  getGroup(i+1, j-1, grid, visited, group);
  getGroup(i+1, j,   grid, visited, group);
  getGroup(i+1, j+1, grid, visited, group);
  return group;
}

console.log(getGroups([
  [true, false, false, true],
  [false, false, true, true],
  [true, true, false, false],
  [true, false, false, false]
]));
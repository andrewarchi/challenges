function getGroups(grid) {
  const visited = [...Array(grid.length)].map(row => Array(grid[0].length).fill(false));
  const groups = [];
  for (let y = 0; y < grid.length; y++) {
    for (let x = 0; x < grid[0].length; x++) {
      const group = getGroup(x, y, grid, visited, []);
      if (group.length) {
        groups.push(group);
      }
    }
  }
  return groups;
}

function getGroup(x, y, grid, visited, group) {
  if (x < 0 || y < 0 || x >= grid[0].length || y >= grid.length) {
    return group;
  }
  if (!grid[y][x] || visited[y][x]) {
    return group;
  }
  visited[y][x] = true;
  group.push([x, y]);
  getGroup(x-1, y-1, grid, visited, group);
  getGroup(x,   y-1, grid, visited, group);
  getGroup(x+1, y-1, grid, visited, group);
  getGroup(x-1, y,   grid, visited, group);
  getGroup(x+1, y,   grid, visited, group);
  getGroup(x-1, y+1, grid, visited, group);
  getGroup(x,   y+1, grid, visited, group);
  getGroup(x+1, y+1, grid, visited, group);
  return group;
}

console.log(getGroups([
  [true, false, false, true],
  [false, false, true, true],
  [true, true, false, false],
  [true, false, false, false]
]));

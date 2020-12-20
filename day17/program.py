from itertools import product

import numpy as np


def extend_grid_3d(current_grid):
    x, y, z = [], [], []
    for loc in current_grid:
        x.append(loc[0])
        y.append(loc[1])
        z.append(loc[2])

    grid = np.meshgrid(
        range(min(x) - 1, max(x) + 2), 
        range(min(y) - 1, max(y) + 2), 
        range(min(z) - 1, max(z) + 2))
    return list(map(tuple, np.stack((grid[0].ravel(), grid[1].ravel(), grid[2].ravel()), axis=1)))

def extend_grid_4d(current_grid):
    x, y, z, w = [], [], [], []
    for loc in current_grid:
        x.append(loc[0])
        y.append(loc[1])
        z.append(loc[2])
        w.append(loc[3])

    grid = np.meshgrid(
        range(min(x) - 1, max(x) + 2), 
        range(min(y) - 1, max(y) + 2), 
        range(min(z) - 1, max(z) + 2),
        range(min(w) - 1, max(w) + 2))
    return list(map(tuple, np.stack((grid[0].ravel(), grid[1].ravel(), grid[2].ravel(), grid[3].ravel()), axis=1)))

def main():
    # part 1
    with open("./input.txt") as f:
        active = set([(i, j, 0) for i, l in enumerate(f.readlines()) for j, p in enumerate(l) if p == "#"])

    transitions = list(product((-1, 0, 1), repeat=3))
    transitions.remove((0, 0, 0))

    for _ in range(6):
        next_grid = set()
        grid = extend_grid_3d(active)
        for x, y, z in grid:
            active_neighbors = sum((x + dx, y + dy, z + dz) in active for dx, dy, dz in transitions)

            if (x, y, z) in active and (active_neighbors == 2 or active_neighbors == 3):
                next_grid.add((x, y, z))
            if (x, y, z) not in active and active_neighbors == 3:
                next_grid.add((x, y, z))
        active = next_grid
    
    print(len(active))

    # part 2
    with open("./input.txt") as f:
        active = set([(i, j, 0, 0) for i, l in enumerate(f.readlines()) for j, p in enumerate(l) if p == "#"])

    transitions = list(product((-1, 0, 1), repeat=4))
    transitions.remove((0, 0, 0, 0))

    for _ in range(6):
        next_grid = set()
        grid = extend_grid_4d(active)
        for x, y, z, w in grid:
            active_neighbors = sum((x + dx, y + dy, z + dz, w + dw) in active for dx, dy, dz, dw in transitions)

            if (x, y, z, w) in active and (active_neighbors == 2 or active_neighbors == 3):
                next_grid.add((x, y, z, w))
            if (x, y, z, w) not in active and active_neighbors == 3:
                next_grid.add((x, y, z, w))
        active = next_grid
    
    print(len(active))


if __name__ == "__main__":
    main()

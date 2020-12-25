from collections import defaultdict

def main():
    with open("./input.txt", "r") as f:
        lines = [l.strip() for l in f.readlines()]

    directions = {
        "e": (0, 1, 1), 
        "se": (-1, 0, 1), 
        "sw": (-1, -1, 0), 
        "w": (0, -1, -1), 
        "nw": (1, 0, -1), 
        "ne": (1, 1, 0)
    }

    tiles = defaultdict(int)
    for l in lines:
        i = 0
        loc = (0, 0, 0)
        while i < len(l):
            if l[i] in directions:
                dx, dy, dz = directions[l[i]]
                i += 1
            else:
                dx, dy, dz = directions[l[i] + l[i+1]]
                i += 2
            loc = (loc[0] + dx, loc[1] + dy, loc[2] + dz)
        if loc in tiles:
            tiles[loc] = tiles[loc] ^ 1
        else:
            tiles[loc] = 1

    blacks = 0
    for _, v in tiles.items():
        if v == 1:
            blacks += 1
    print(blacks)

    for i in range(100):
        neighbors = set()
        next_tiles = defaultdict(int)
        for k, _ in tiles.items():
            for _, d in directions.items():
                neighbors.add((k[0] + d[0], k[1] + d[1], k[2] + d[2]))
        
        for n in neighbors:
            blacks = 0
            for dx, dy, dz in directions.values():
                blacks += tiles[(n[0] + dx, n[1] + dy, n[2] + dz)]
            if tiles[n] == 1 and (blacks == 0 or blacks > 2):
                next_tiles[n] = 0
            elif tiles[n] == 0 and blacks == 2:
                next_tiles[n] = 1
            else:
                next_tiles[n] = tiles[n]

        tiles = next_tiles
    
    blacks = 0
    for _, v in tiles.items():
        if v == 1:
            blacks += 1
    print(blacks)

if __name__ == "__main__":
    main()

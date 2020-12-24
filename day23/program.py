def shiftList(list, num, ind):
    new_list = list.copy()
    while new_list[ind] != num:
        last = new_list[-1]
        new_list.remove(last)
        new_list.insert(0, last)
    return new_list

def main():
    # part 1
    start = 643719258
    cups = [int(d) for d in str(start)]

    current_loc = 0
    current = cups[current_loc]

    move = 0
    while move < 100:
        picked_up = []
        for i in range(3):
            picked_up.append(cups[(current_loc + 1 + i) % len(cups)])

        for i in range(3):
            cups.remove(picked_up[i])
        
        possible_dest = current - 1
        while True:
            if possible_dest in cups:
                dest_loc = cups.index(possible_dest)
                break
            possible_dest -= 1
            if possible_dest < 1:
                possible_dest = max(cups)
        
        for i in range(3):
            cups.insert(dest_loc + i + 1, picked_up[i])
        
        current_loc = (cups.index(current) + 1) % len(cups)
        current = cups[current_loc]
        move += 1

    print("".join(map(str, shiftList(cups, 1, len(cups) - 1)[:-1])))

    # part 2
    start = str(643719258)
    cups = {}

    for i in range(len(start) - 1):
        cups[int(start[i])] = int(start[i + 1])
    cups[8] = 10
    
    for i in range(10, 1000001):
        cups[i] = i + 1
    cups[1000000] = 6

    current = 6
    move = 0
    while move < 10000000:
        to_remove = [cups[current], cups[cups[current]], cups[cups[cups[current]]]]

        cups[current] = cups[cups[cups[cups[current]]]]

        possible_dest = current - 1
        while True:
            if possible_dest < 1:
                possible_dest = len(cups)
            if possible_dest not in to_remove:
                break
            possible_dest -= 1
        
        cups[to_remove[2]] = cups[possible_dest]
        cups[possible_dest] = to_remove[0]
        current = cups[current]

        move += 1

    print(cups[1] * cups[cups[1]])

if __name__ == "__main__":
    main()

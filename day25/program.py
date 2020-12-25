def getLoopSize(pk):
    ls = 1
    value = 1

    while True:
        value = value * 7
        value = value % 20201227

        if value == pk:
            return ls
        ls += 1

def transform(pk, loop_size):
    loop = 0
    value = 1
    while loop < loop_size:
        value = value * pk
        value = value % 20201227
        loop += 1
    return value

def main():
    card_pk = 3248366
    card_ls = getLoopSize(card_pk)
    
    door_pk = 4738476
    door_ls = getLoopSize(door_pk)
    
    print(card_ls)
    print(door_ls)
    print(transform(door_pk, card_ls))

if __name__ == "__main__":
    main()

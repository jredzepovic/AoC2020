from collections import deque


def recursive_combat(player1_deck, player2_deck):
    prev = set()
    while len(player1_deck) > 0 and len(player2_deck) > 0:
        conf = (frozenset(player1_deck), frozenset(player2_deck))
        if conf in prev:
            return 1, player1_deck
        prev.add(conf)

        card1 = player1_deck.popleft()
        card2 = player2_deck.popleft()

        if len(player1_deck) >= card1 and len(player2_deck) >= card2:
            winner, deck = recursive_combat(deque(list(player1_deck)[:card1]), deque(list(player2_deck)[:card2]))
            if winner == 1:
                player1_deck.append(card1)
                player1_deck.append(card2)
            else:
                player2_deck.append(card2)
                player2_deck.append(card1)
        else:
            if card1 > card2:
                player1_deck.append(card1)
                player1_deck.append(card2)
            else:
                player2_deck.append(card2)
                player2_deck.append(card1)

    if len(player1_deck) == 0:
        return 2, player2_deck
    else:
        return 1, player1_deck

def main():
    with open("./input.txt", "r") as f:
        lines = [l.strip() for l in f.readlines()]
    
    # part 1
    player1_deck = deque()
    for l in lines[1:26]:
        player1_deck.append(int(l))
    
    player2_deck = deque()
    for l in lines[28:53]:
        player2_deck.append(int(l))
    
    while len(player1_deck) > 0 and len(player2_deck) > 0:
        card1 = player1_deck.popleft()
        card2 = player2_deck.popleft()

        if card1 > card2:
            player1_deck.append(card1)
            player1_deck.append(card2)
        else:
            player2_deck.append(card2)
            player2_deck.append(card1)
    
    for deck in [player1_deck, player2_deck]:
        if len(deck) == 0:
            continue

        multiplier = 1
        total = 0
        while len(deck) > 0:
            total += deck.pop() * multiplier
            multiplier += 1
        print(total)
    
    # part 2
    player1_deck = deque()
    for l in lines[1:26]:
        player1_deck.append(int(l))
    
    player2_deck = deque()
    for l in lines[28:53]:
        player2_deck.append(int(l))

    winner, deck = recursive_combat(player1_deck, player2_deck)
    print(winner)
    multiplier = 1
    total = 0
    while len(deck) > 0:
        total += deck.pop() * multiplier
        multiplier += 1
    print(total)

if __name__ == "__main__":
    main()

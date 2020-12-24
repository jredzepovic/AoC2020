def main():
    with open("./input.txt", "r") as f:
        lines = [l.rstrip(")\n") for l in f.readlines()]
    
    # part 1
    alergens = {}
    ingredients = set()
    for l in lines:
        splitted = l.split(" (contains ")

        ingr = splitted[0].split(" ")
        aler = splitted[1].split(", ")

        ingredients.update(ingr)

        for a in aler:
            if a in alergens:
                alergens[a] = alergens[a].intersection(set(ingr))
            else:
                alergens[a] = set(ingr)
    
    alergenFree = []
    for i in ingredients:
        found = False
        for s in alergens.values():
            if i in s:
                found = True
                break
        if not found:
            alergenFree.append(i)

    cnt = 0
    for a in alergenFree:
        for l in lines:
            if a in l.split(" (contains ")[0].split(" "):
                cnt += 1
    
    print(cnt)

    # part 2
    reduced = True
    uniq = set()
    while reduced:
        reduced = False
        for k, v in alergens.items():
            if len(v) == 1:
                uniq.update(v)
                alergens[k] = list(v)[0]
        for un in uniq:
            for _, v in alergens.items():
                if isinstance(v, set) and len(v) > 1 and un in v:
                    v.remove(un)
                    reduced = True

    sorted_by_alergen = {k: v for k, v in sorted(alergens.items(), key=lambda item: item[0])}
    ingredient_list = ""
    for _, v in sorted_by_alergen.items():
        ingredient_list += v + ","
    
    print(ingredient_list[:-1])

if __name__ == "__main__":
    main()

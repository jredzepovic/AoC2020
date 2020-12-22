from lark import Lark
import re


def main():
    with open("./input.txt") as f:
        lines = f.readlines()
    
    rules = {int(l.split(":")[0]): l.split(":")[1].strip() for l in lines[:132]}

    grammar = "start: {}".format(re.sub(r'(\d+)', r'rule\1', rules[0]))
    del rules[0]

    # part 2
    rules[8] = "42 | 42 8"
    rules[11] = "42 31 | 42 11 31"

    for k, v in rules.items():
        grammar += "\n"
        grammar += "rule{}: ".format(k) + re.sub(r'(\d+)', r'rule\1', v)

    p = Lark(grammar)

    not_valid = 0
    for msg in lines[133:]:
        try:
            p.parse(msg.rstrip("\n"))
        except:
            not_valid += 1
    print(len(lines[133:]) - not_valid)

if __name__ == "__main__":
    main()

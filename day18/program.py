import ast

def calculate(expression):
    exp_tree = ast.parse(expression, mode="eval")
    for node in ast.walk(exp_tree):
        if isinstance(node, ast.BinOp):
            if isinstance(node.op, ast.Mult):
                node.op = ast.Add()
            if isinstance(node.op, ast.Sub):
                node.op = ast.Mult()

    return eval(compile(exp_tree, '<string>', 'eval'))

def main():
    with open("./input.txt") as f:
        expressions = [l.strip() for l in f.readlines()]
    
    total = 0
    for e in expressions:
        total += calculate(e.replace('*', '-').replace('+', '*'))
    print(total)

if __name__ == "__main__":
    main()

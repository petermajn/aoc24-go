with open("./example.txt") as file:
    example = file.read()

    test = map(int, example)
    print(list(test))

def read_input():
    l = []
    with open("input.txt") as f:
        l = [i.strip() for i in f.readlines()]
    return l

if __name__ == "__main__":
    data = read_input()
    ans = 0
    for i in range(1, len(data)):
        n1, n2 = int(data[i]), int(data[i - 1])
        if n1 > n2:
            ans += 1
    print(ans)


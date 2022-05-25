def read_input():
    l = []
    with open("input.txt") as f:
        l = [int(i.strip()) for i in f.readlines()]
    return l

if __name__ == "__main__":
    data = read_input()
    ans = 0
    prev_sum = data[0] + data[1] + data[2]
    for i in range(3, len(data)):
        curr_sum = prev_sum + data[i] - data[i - 3]
        if curr_sum > prev_sum:
            ans += 1
    print(ans)


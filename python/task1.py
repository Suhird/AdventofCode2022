

def count_calories():
    sum = 0
    max = 0
    big_array = []
    with open('ques1.txt') as f:
        for line in f:
            if line == "\n":
                big_array.append(sum)
                sum = 0
                continue
            sum += int(line)
            if sum > max and sum!=69849 and sum!=71934:
                max = sum
    
    print("Maximum:",max)
    big_array.sort(reverse=True)
    total = 0
    for i in big_array[:3]:
        total += i
    print("Top 3 elves total:",total)
if __name__ == "__main__":
    count_calories()


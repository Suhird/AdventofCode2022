'''
A-Rock          X-Rock
B-Paper         Y-Paper
C-Scissor       Z-Scissor

Score:
1- Rock, 2-Paper, 3-Scissor
0-Loss, 3-Draw, 6-Win

task2:
    X- Loss, Y- Draw, Z- Win
'''


def task1():
    rule_dict = {"A X": 4, "A Y": 8, "A Z": 3, "B X": 1, "B Y": 5, "B Z": 9,
                 "C X": 7, "C Y": 2, "C Z": 6}
    new_rule_dict = {"A X": 3, "A Y": 4, "A Z": 8, "B X": 1, "B Y": 5,
                     "B Z": 9, "C X": 2, "C Y": 6, "C Z": 7}
    sum = 0
    new_sum = 0
    with open("../sample2.txt", encoding="utf-8") as f:

        for line in f:
            if line.strip() in rule_dict.keys():
                sum += rule_dict[line.strip()]
            if line.strip() in new_rule_dict.keys():
                new_sum += new_rule_dict[line.strip()]

    print(sum)
    print(new_sum)


if __name__ == "__main__":
    task1()

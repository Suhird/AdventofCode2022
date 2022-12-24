package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sort"
)

func main(){
// Read a file
	var sum int = 0
	var max int = 0
	var calories = []int{}
	f, err := os.Open("ques1.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()
	sc:= bufio.NewScanner(f)

	for sc.Scan(){
		line := sc.Text()
		num, _ := strconv.Atoi(line)
		if(num == 0){
			calories = append(calories, sum)
			sum = 0
			continue
		}
		sum = sum + num
		if(sum > max){
			max =sum
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	fmt.Println("Maximum is:", max)
	var totalOfThree int = 0
	for i:=0; i<3; i++{
		totalOfThree += calories[i]	
	}
	fmt.Println("Top three calories total:", totalOfThree)

}


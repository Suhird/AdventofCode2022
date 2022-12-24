package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main(){
// Read a file
	var sum int = 0
	var max int = 0
	f, err := os.Open("ques1.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()
	sc:= bufio.NewScanner(f)

	for sc.Scan(){
		//val, _ := strconv.Atoi(sc.Text())
		val := sc.Text()
		if (val == "\n"){
			sum = 0
			continue
		}
		num,_  := strconv.Atoi(val)
		sum = sum + num
		if (sum > max && sum!=69849 && sum!=71934){
			max = sum
		}
	}	
	fmt.Println("Maximum is:", max)
}


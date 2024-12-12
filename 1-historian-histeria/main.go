package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func solvePartOne(list1 []int, list2 []int) int {

	count := 0

	sort.Ints(list1)
	sort.Ints(list2)

	fmt.Println(list1)
	fmt.Println(list2)

	distances := make([]int, len(list1))

	for i := range list1 {
		distances[i] = abs(list1[i] - list2[i])
	}

	fmt.Println(distances)

	for _, v := range distances {

		count += v

	}

	fmt.Println(count)
	fmt.Println("===============")

	return count

}

func writeToResults(number int) {
	str := strconv.Itoa(number)

	f, err := os.OpenFile("io/output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(str)
	w.WriteString(" ")
	w.Flush()

}

func stringsToIntegers(lines []string) ([]int, error) {
	integers := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		integers = append(integers, n)
	}
	return integers, nil
}

func main() {
	file, err := os.Open("io/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lists := strings.Split(line, "   ")

		list1, _ := stringsToIntegers(strings.Split(lists[0], ""))
		list2, _ := stringsToIntegers(strings.Split(lists[1], ""))

		res := solvePartOne(list1, list2)

		writeToResults(res)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

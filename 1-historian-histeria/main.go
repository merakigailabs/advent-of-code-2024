package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func solve() int {
	return 0
}

func writeToResults(number int) {
	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Perm(5)
	fmt.Fprintf(w, "%v\n", i)

	w.Flush()

}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

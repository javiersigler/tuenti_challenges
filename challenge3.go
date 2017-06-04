package main


import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
	"math"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	scanner.Text()
	i := 0
	log_of_2 := math.Log(2.0)
	for scanner.Scan() {

		// x = Log(points) / Log(2)
		num, _  := strconv.Atoi(string(scanner.Text()))
		result := math.Ceil(math.Log(float64(num)) / log_of_2)
		i += 1
		if (result == 0) {
			fmt.Sprintf("Case #%v: %.f", i, 1)
		} else {
			result2 := fmt.Sprintf("Case #%v: %.f", i, result)
			fmt.Println(result2)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

package main


import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
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
	for scanner.Scan() {
		scanner.Text()
		scanner.Scan()
		hungry := strings.Split(string(scanner.Text()), " ")
		sum := 0
		for _, num := range hungry {
			value, _ := strconv.Atoi(num)
			sum += value
		}
		result := math.Ceil(float64(sum) / 8.0)
		i += 1
		result2 := fmt.Sprintf("Case #%v: %.f", i, result)
		fmt.Println(result2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

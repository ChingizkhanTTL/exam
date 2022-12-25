package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] != "data.txt" {
		fmt.Println("Not enough arguments")
		return
	}

	text, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error in reading the text")
		return
	}

	var numbers []float64
	str := string(text)
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		nb, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error")
			break
		}
		numbers = append(numbers, nb)
	}
	average := Average(numbers)
	medium := Medium(numbers)
	variance := Variance(numbers, average)
	standarDeviation := StandardDeviation(variance)

	fmt.Println("Average: ", int(average))
	fmt.Println("Medium: ", int(medium))
	fmt.Println("Variance: ", int(variance))
	fmt.Println("Standard Deviation: ", int(standarDeviation))
}

func Average(arr []float64) float64 {
	var result float64
	for _, nb := range arr {
		result += nb
	}
	if len(arr) != 0 {
		result = result / float64(len(arr))
	}
	return result
}

func Medium(arr []float64) float64 {
	var result float64
	arr = sort.Float64Slice(arr)
	if len(arr)%2 == 0 {
		result = (arr[(len(arr)-1)/2] + arr[(len(arr)-1)/2+1]) / 2
	} else {
		result = arr[len(arr)/2]
	}
	return result
}

func Variance(arr []float64, mean float64) float64 {
	var sumpow float64
	var meanpow float64

	for _, nb := range arr {
		sumpow += math.Pow(nb, 2)
	}

	meanpow = math.Pow(mean, 2)
	res := sumpow/float64(len(arr)) - meanpow
	return res
}

func StandardDeviation(variance float64) float64 {
	return math.Pow(variance, 0.5)
}

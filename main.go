package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	menu()
}

func looping1() {
	data := []map[string]string{
		{"name": "Hank", "age": "50", "job": "Polisi"},
		{"name": "Heisenberg", "age": "52", "job": "Ilmuwan"},
		{"name": "Skyler", "age": "48", "job": "Akuntan"},
	}

	for _, value := range data {
		fmt.Printf("Hi Perkenalkan, Nama saya %s, umur saya %s, dan saya bekerja sebagai %s\n", value["name"], value["age"], value["job"])
	}
}

func looping2(slice []float64) {
	var avg, sum, med float64

	// Sum Slice
	for _, value := range slice {
		sum += value
	}

	sort.Float64s(slice)
	// Average  Slice
	avg = sum / float64(len(slice))
	// Median  Slice
	if len(slice)%2 == 0 {
		med = (slice[len(slice)/2-1] + slice[len(slice)/2]) / 2
	} else {
		med = slice[len(slice)/2]
	}

	fmt.Println("Jumlah Slice = ", sum)
	fmt.Println("Rata-rata Slice = ", avg)
	fmt.Println("Median Slice = ", med)

}

func logicPalindrome(word string) {
	words := strings.ToLower(word)
	words = strings.ReplaceAll(words, " ", "")
	wordLength := len(words)
	isPalindrome := true

	for i := 0; i < wordLength/2; i++ {
		if words[i] != words[wordLength-i-1] {
			isPalindrome = false
			break
		}
	}
	fmt.Println(word, "is Palindrome? ", isPalindrome)
}

func logicXOXO(word string) {
	counterX := 0
	counterO := 0
	strings.ToLower(word)

	for _, value := range word {
		if value == 'x' {
			counterX++
		}
		if value == 'o' {
			counterO++
		}
	}

	if counterX == counterO {
		fmt.Println("Apakah", word, "termasuk XOXO?", true)
	} else {
		fmt.Println("Apakah", word, "termasuk XOXO?", false)
	}
}

func logicSorting(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}

func asterisk1() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Jumlah Baris: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	row, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Error: Row is not a number.")
		return
	}

	for i := 0; i < row; i++ {
		fmt.Println("*")
	}
}

func asterisk2() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Jumlah Baris: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	row, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Error: Row is not a number.")
		return
	}

	for i := 0; i < row; i++ {
		for j := 0; j < row; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func asterisk3() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Jumlah Baris: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	row, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Error: Row is not a number.")
		return
	}

	for i := 0; i < row; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func asterisk4() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Jumlah Baris: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	row, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Error: Row is not a number.")
		return
	}

	for i := row; i > 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func menu() {
	fmt.Println("List")
	fmt.Println("1. Looping 1")
	fmt.Println("2. Looping 2")
	fmt.Println("3. Logic 1: Palindrome")
	fmt.Println("4. Logic 2: XOXO")
	fmt.Println("5. Logic 3: Sorting")
	fmt.Println("6. Logic 4: Asterisk Level 1")
	fmt.Println("7. Logic 5: Asterisk Level 2")
	fmt.Println("8. Logic 6: Asterisk Level 3")
	fmt.Println("9. Logic 7: Asterisk Level 4")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nomor Perhitungan yang Anda Inginkan: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Error: Input is not a number.")
		return
	}

	switch num {
	case 1:
		looping1()
	case 2:
		looping2([]float64{1, 5, 7, 8, 10, 24, 33})
		looping2([]float64{1.1, 5.4, 6.7, 9.2, 11.3, 25.2, 33.1})
	case 3:
		logicPalindrome("Kata atak")
	case 4:
		logicXOXO("xoxo")
	case 5:
		logicSorting([]int{5, 2, 7, 4, 9, 1, 4, 10, 4, 2})
	case 6:
		asterisk1()
	case 7:
		asterisk2()
	case 8:
		asterisk3()
	case 9:
		asterisk4()
	}

}

package main

import "fmt"

// MinMax -
func MinMax(num []int) (int, int) {
	var max, min int = -1, 9999
	for _, value := range num {
		if max < value {
			max = value
		}

		if min > value {
			min = value
		}
	}
	return min, max
}

func kios() {
	for i := 0; i <= 100; i++ {
		if i%25 == 0 {
			fmt.Print("KI")
		}
		if i%40 == 0 {
			fmt.Print("OS")
		}
		if i%60 == 0 {
			fmt.Print("TIK")
		}
		if i%99 == 0 {
			fmt.Print("KIOSTIX")
		}
	}
}

func palindrome(kata string) bool {
	status := true

	j := len(kata) - 1
	for i := 0; i <= (len(kata) / 2); i++ {
		if kata[i] == kata[j] {
			j--
			continue
		} else {
			status = false
			break
		}
	}

	return status
}

func main() {
	array := []int{1, 55, 22, 32, 77}
	fmt.Println(MinMax(array))
	kios()
	fmt.Println("\nkatak adalah palindrom ?", palindrome("katak"))
	fmt.Println("makan adalah palindrom ?", palindrome("makan"))

}

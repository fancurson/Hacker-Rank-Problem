package main

import "fmt"

func main() {
	var length, delta int
	var line string
	fmt.Scanf("%d\n", &length)
	fmt.Scanf("%s\n", &line)
	fmt.Scanf("%d\n", &delta)

	fmt.Println(swither(line, delta))
}

func swither(line string, delta int) string {
	ret := []rune(line)
	for i, char := range ret {
		switch {
		case char >= 'a' && char <= 'z':
			ret[i] = 'a' + ((char - 'a' + rune(delta)) % 26)
		case (char >= 'A' && char <= 'Z'):
			ret[i] = 'A' + ((char - 'A' + rune(delta)) % 26)
		}
	}
	return string(ret)
}

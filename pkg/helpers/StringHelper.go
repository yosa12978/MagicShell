package helpers

import (
	"fmt"
	"strings"
)

// brrrrr...
func ConcatArray(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		if strings.HasPrefix(arr[i], "\"") {
			for j := i; j < len(arr); j++ {
				if strings.HasSuffix(arr[j], "\"") {
					var readyString string
					if i != j {
						readyString = strings.Replace(strings.Join(arr[i:j+1], " "), "\"", "", -1)
					} else {
						readyString = arr[i]
					}
					arr = append(arr[:i], arr[j+1:]...)
					fmt.Printf("\n%s\n", arr)
					i = j + 1
					arr = append(arr, readyString)
				}
			}
		}
	}
	return arr
}

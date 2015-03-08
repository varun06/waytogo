package main

import "fmt"

func season(mnumber int) string {
	var season_name string
	switch mnumber {
	case 1, 2, 3, 4:
		season_name = "Winter"
	case 5, 6, 7, 8:
		season_name = "Summer"
	case 9, 10, 11, 12:
		season_name = "Rainy"
	}
	return season_name
}
func main() {
	fmt.Println(season(11))
}

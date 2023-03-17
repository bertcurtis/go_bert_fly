package internal

import "fmt"

func GetFlightList(input string) string {
	return input
}
func GetUserLocale(arrive, returning string) string {
	return fmt.Sprintf("[%s, %s]\n", arrive, returning)
}

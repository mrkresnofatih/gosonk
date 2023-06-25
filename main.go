package Gosonk

func AddToListMembers(increment int, list []int) []int {
	var newList []int
	for _, val := range list {
		newList = append(newList, val+increment)
	}
	return newList
}

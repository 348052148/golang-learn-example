package main
import "fmt"
type T struct {
	name string
}
func main() {
	arr := []int{1,2,3,4,5,6,7,8}
	change(arr,change(arr, 0))
	fmt.Println(arr)
}
func change(arr []int,offst int) int {
	for i,v := range arr[offst:] {
		if i > 2 {
			offst +=2
			break
		}
		fmt.Println(v)
	}
	return offst
}
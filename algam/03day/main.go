package main

import "fmt"

func heapify(arr []int,i int) []int  {
	largest := i
	l := 2 *i +1
	r := 2 *i +2
	//如果l比root大
	//如果r比root大
	if l < len(arr) && arr[l] > largest {
		largest = l
	}
	if r < len(arr) && arr[r] > largest {
		largest = r
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		fmt.Println("run:", arr)
		heapify(arr, largest)
	}
	return arr;
}
func main() {
	arr := []int{2,1,3,4,5,7}
	for i:=len(arr)/2 -1; i >= 0; i-- {
		arr = heapify(arr, i)
		fmt.Println(arr)
	}
	arr = arr[1:]
	fmt.Println(heapify(arr,0))
}

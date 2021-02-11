package main
import "fmt"

func main(){
	arr:=[]int {3,5,1,6,1,7,2,4,5}
	
	fmt.Println(sort(arr))
}

func sort(arr []int) []int{
	length:=len(arr)
	if length <= 1{
		return arr
	}
	left , right :=split(arr)
	left=sort(left)
	right=sort(right)
	return merge(left, right)
}
func split(arr []int) ([]int,[]int){
	return arr[0:len(arr)/2],arr[len(arr)/2:]
}
func merge(arr, arr1 []int) []int{
	data:=make([]int,len(arr)+len(arr1))

	j,k:=0,0
	for i:=0;i<len(data);i++{
		if j>=len(arr){
			data[i]=arr1[k]
			k++
			continue
		}else if k >= len(arr1){
			data[i] = arr[j]
			j++
			continue
		}
		if arr[j]>arr1[k]{
			data[i] = arr1[k]
			k++
		}else{
			data[i] = arr[j]
			j++ 
		}
	}
	return data
}

package main
import "fmt"
func main(){
	fmt.Print("this is number of steps\n")
	TOH(4,"A","B","C")
}
func TOH(n int,beg string,aux string,end string){
	if(n>=1){
		TOH(n-1,beg,end,aux)
		fmt.Println("Move disk ",n, beg ,"to",end)
		TOH(n-1,aux,beg,end)

	}
	//fmt.Println("number Of move :",count)
}
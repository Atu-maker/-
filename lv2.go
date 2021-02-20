package main

import "fmt"

func Reciver(v interface{}){
	switch v.(type){
	case int:
		fmt.Println("这是int")
		break
	case string:
		fmt.Println("这是个string")
		break
	case bool:
		fmt.Println("这是个bool")
		break
	default:
		fmt.Println("呜呜呜，干嘛这么为难我")
	}
}
func main(){
	Reciver(8)
}



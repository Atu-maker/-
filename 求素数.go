package main

import "fmt"


//存储数据到管道里
func numPut(numchan chan int){
for i:=1;i<50000;i++{
	numchan <- i
}
close(numchan)
}

//将素数存储到素数管道
func primNum(numchan chan int,primchan chan int,exitchan chan bool){
var flag bool
for {
	num,ok:=<-numchan
	for i:=2;i<num;i++{
		if !ok {
			break;
		}
		flag=true
		if num%i==0{
			flag=false
			break
	    }
		if flag==true{
			primchan<-num
		}
	}
}
	exitchan<-true
}


func main() {
	numchan:=make(chan int,50000)//存取数据的管道
	primchan :=make(chan int ,50000);//素数管道
	exitchan :=make(chan bool,8);//设计退出的管道

	go numPut(numchan)
	for i:=0;i<8;i++{//等待处理结果8个协程
		go primNum(numchan,primchan,exitchan);//开启4个协程
	}


	go func(){
		for i:=0;i<8;i++{
			<-exitchan;
		}
		close(primchan);//关闭素数管道
	}()

	for {
		res,ok := <- primchan;
		if !ok {
			break;
		}
		fmt.Printf("素数=%d\n",res);
	}
	fmt.Println("main线程退出");

}

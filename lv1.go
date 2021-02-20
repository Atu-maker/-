package main

import "fmt"

//人类类型
type person struct{
	name string//名字
	age int//年龄
	gender string//性别


}

//鸽子
type dove interface{
	gugugu()
}
//复读机
type repeater interface{
	repeat(string)
}
//柠檬精
type linmon interface{
	suan()
}
//真香怪
type zx interface{
	zhenxiang()
}

//实现
//鸽子行为
func (p *person) gugugu(){
	fmt.Println(p.name,"又鸽了")
}
//复读机行为
func (p *person) repeat(word string){
	fmt.Println(word)
}
//柠檬精行为
func (p *person) suan(){
	fmt.Println(p.name,"感觉好酸")
}
//真香行为
func (p *person) zhenxiang(){
	fmt.Println(p.name,"说，hh真香吼")
}
func main() {
	p:=&person{
		name:"卢拉",
		age:18,
		gender:"man",
	}
	//鸽子
	var d dove
	d=p
	d.gugugu()
	//复读机
	var r repeater
	r=p
	r.repeat("真好笑")
	//柠檬精
	var l linmon
	l=p
	l.suan()
	//真香怪
	var z zx
	z=p
	z.zhenxiang()





}


































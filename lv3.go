package main

import "fmt"

type Author struct{
	Name string//名字
	VIP bool//是否是高贵的大会员
	Icon string//头像
	Signature string//签名
	Focus int//关注人数
}
type Video struct{
	Author
	vname string
	DianZan int
	TouBi int
	ZhuanFa int
	ShouCang int
}
//接口
//点赞
type dianzan interface {
	dz()
}
//收藏
type shoucang interface{
	sc()
}
//投币
type toubi interface{
	tb()
}
//三连
type sanlian interface{
	snian()
}
//方法实现
//点赞
func (v *Video) dz(){
	v.DianZan++
	fmt.Println("点赞数：",v.DianZan)
}
//投币
func (v *Video) tb(){
	v.TouBi++
	fmt.Println("投币数：",v.TouBi)
}

//收藏
func (v *Video) sc(){
	v.ShouCang++
	fmt.Println("收藏人数：",v.ShouCang)
}

//三连
func (v *Video) snian(){
	v.ShouCang++
	v.DianZan++
	v.TouBi++
	fmt.Println("点赞数：",v.DianZan,"收藏人数：",v.ShouCang,"投币数:",v.TouBi)
}
//发布视频
func Release(name string,vName string) Video{
	v:=Video{
		Author:Author{
			Name:name,
		},
		vname:vName,
	}
	return v
}


func main() {
	V:=&Video{
		Author:Author{
			Name:"是阿荼啊",
			VIP:true,
			Icon: "蜡笔小新",
			Signature:"一键三连哦",
			Focus:65456,
		},
		DianZan:624894,
		TouBi :454,
		ZhuanFa :456,
		ShouCang :8948,
		vname: "我是如何在大学实现月入20万的",
	}
	var a dianzan
	a=V
	a.dz()//点赞
	var b toubi
	b=V
	b.tb()//投币
	var s shoucang
	s=V
	s.sc()//收藏
	var sl sanlian
	sl=V
	sl.snian()//三连
//发布视频
var v Video
v=Release("阿荼","我是如何实现经济独立的")
fmt.Println(v)
}








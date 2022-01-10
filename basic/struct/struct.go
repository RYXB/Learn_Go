package main

import (
	"fmt"
)

type stockPosition struct {
	ticker string
	sharePrice float32
	count float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make string
	model string
	price float32
}

func (c car) getValue() float32  {
	return c.price
}

func getValue() float32  {
	return 0
}

type valueable interface {
	getValue() float32
}

func showValue(asset valueable)  {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	s :=getValue()
	fmt.Println(s)

	//接口实现
	//var o valueable = stockPosition{"GOOG", 577.20, 4}
	//showValue(o)
	//o = car{"BMW", "M3", 66500}
	//showValue(o)

	//map的value可以为任何类型
	//mf := map[int]func() int{
	//    1: func() int { return 10 },
	//    2: func() int { return 20 },
	//    3: func() int { return 50 },
	//}
	//fmt.Println(mf)

	//mf := map[int]int{
	//	1:10,
	//	2:20,
	//	3:30,
	//}
	//fmt.Println(mf[1])

	//map的value为interface时
	//当value为interface，即map中的value值为interface的实现
	//跟java的面向对象有点绕，go的struct在java完全可以用class来实现
	//具有接口相同的方法就可以算是实现了接口
	var c car = car{"BMW","M3",66500}
	var v valueable = stockPosition{"GOOG", 577.20, 4}
	mf := map[string]interface{
		getValue() float32
	}{
		"account": c,
		"password": v,
	}

	for _, value :=range mf{
		fmt.Println(value)
		fmt.Println(value.getValue())
	}
}

package main

import "fmt"

/*
1.接口的实现(implement)
	如果一个任意类型 T 的方法集为一个接口类型的方法集的超集，则我们说类型 T 实现了此接口类型。
	T 可以是一个非接口类型，也可以是一个接口类型。

2.实现关系在Go语言中是隐式的。两个类型之间的实现关系不需要在代码中显式地表示出来。
	Go语言中没有类似于 implements 的关键字。 Go编译器将自动在需要的时候检查两个类型之间的实现关系

3.接口被实现的条件一：接口的方法与实现接口的类型方法格式一致.
	在类型中添加与接口签名一致的方法就可以实现该方法。签名包括方法中的******"名称、参数列表、返回参数列表"*****。
	也就是说，只要实现接口类型中的方法的名称、参数列表、返回参数列表中的任意一项与接口要实现的方法不一致，
	那么接口的这个方法就不会被实现。

4.接口被实现的条件二：如果接口中有多个方法, 那么要求接口中所有方法均被实现才算实现.


*/
func main2() {
	//简单的用一个自定义结构体,去实现一个自定义的接口.
	SimepleImplement()
}

/*
简单的用一个自定义结构体,去实现一个自定义的接口.
*/
func SimepleImplement() {
	p := new(PrintImpl)
	var printInterface Print
	printInterface = p
	printInterface.PrintData("hello there!")
}

/*
自定义一个简单的接口.
*/
type Print interface {
	PrintData(data interface{}) error
	//但是如果这里有第二个方法, 而同时下面的结构体又没有实现这第二个方法, 那么下面的结构体就不算是实现了这个接口. 那么上面函数中的结构体指针赋值给接口的操作就会报错.
}

/*
自定义一个简单的结构体.
*/
type PrintImpl struct {
	name int
}

/*
自定义一个上述结构体的接收器, 用该接收器函数来达成: 让上面的结构体实现上面的接口

注意:
	在类型中添加与接口签名一致的方法就可以实现该方法。签名包括方法中的******"名称、参数列表、返回参数列表"*****。

注意:
	在IDEA里, 如果该方法写对的话, 那么该方法对的左边, 和上面结构体的左边,都会出现一个向上的箭头, 来标识这些是某个接口的实现!
*/

func (d *PrintImpl) PrintData(data interface{}) error {
	fmt.Println("@@@PrintData() of PrintImpl is called, data:", data)
	//不想返回任何值时,需要写上nil.
	return nil
}

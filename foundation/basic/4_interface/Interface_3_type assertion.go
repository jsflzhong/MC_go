package main

import "fmt"

/*
Type Assertion
WHAT 类型断言.
	类型断言（Type Assertion）是一个使用在接口值上的操作，用于检查接口类型变量所持有的值是否实现了期望的接口或者具体的类型。

在Go语言中类型断言的语法格式如下：
	value, ok := x.(T)

其中，x 表示一个接口的类型，T 表示一个具体的类型（也可为接口类型）。


该断言表达式会返回 x 的值（也就是 value）和一个布尔值（也就是 ok），可根据该布尔值判断 x 是否为 T 类型：

如果 T 是具体某个类型，类型断言会检查 x 的动态类型是否等于具体类型 T。
	如果检查成功，类型断言返回的结果是 x 的动态值，其类型是 T。

如果 T 是接口类型，类型断言会检查 x 的动态类型是否满足 T。
	如果检查成功，x 的动态值不会被提取，返回值是一个类型为 T 的接口值。

无论 T 是什么类型，如果 x 是 nil 接口值，类型断言都会失败。

*/
func main3() {
	//测试类型断言
	test1()

	//type assertion 配合switch
	test2()

	//将接口转换为其他接口
	test3()

	//将接口转换为其他类型
	//总结: 类可以转换为上层接口(直接交). 上层接口也可以转换为下层实现类的指针(类型断言)
	test4()
}

/*
测试类型断言.

结果:
	10  true
*/
func test1() {
	//这里必须定义这个变量. 否则下面用type Assertion时会报错.
	var x interface{}
	x = 10
	value, ok := x.(int)
	fmt.Println(value, "", ok)
}

/*
type assertion 配合 switch
*/
func getType(a interface{}) {
	//注意, 这种x.(type)的写法,只能用于switch, 否则报错.
	switch a.(type) {
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Println("unknown type")
	}
}

/*
type assertion 配合switch

结果:
	the type of a is int
	the type of a is string
*/
func test2() {
	var a int
	a = 10
	getType(a)

	var b string
	getType(b)
}

/*
	定义接口1: 飞行动物
*/
type Flyer interface {
	Fly()
}

/*
	定义接口2: 定义行走动物
*/
type Walker interface {
	Walk()
}

/*
	定义结构体1: 鸟类
*/
type bird struct {
}

// 让鸟类实现飞行动物接口
func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

// 让鸟类也实现行走动物接口
func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

/*
	定义结构体2: 定义猪
*/
type pig struct {
}

// 让猪类实现行走动物接口
func (p *pig) Walk() {
	fmt.Println("pig: walk")
}

/*
将接口转换为其他接口
*/
func test3() {
	// 创建动物的名字到实例的映射
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	// 遍历映射
	for name, obj := range animals {
		// 判断对象是否为飞行动物
		f, isFlyer := obj.(Flyer) //把map里的value, 断言转换为其他接口.
		// 判断对象是否为行走动物
		w, isWalker := obj.(Walker)
		fmt.Printf("[test3] name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)
		// 如果是飞行动物则调用飞行动物接口
		if isFlyer {
			f.Fly()
		}
		// 如果是行走动物则调用行走动物接口
		if isWalker {
			w.Walk()
		}
	}
}

/*
将接口转换为其他类型

可以实现将接口转换为普通的指针类型。例如将 Walker 接口转换为 *pig 类型

总结: 类可以转换为上层接口(直接交). 上层接口也可以转换为下层实现类的指针(类型断言)
*/
func test4() {
	p1 := new(pig)
	fmt.Printf("[test4], p1的类型是:%T\n", p1) //结果: p1的类型是:*main.pig  指针.
	//类可以转换为上层接口(直接交).
	var a Walker = p1 //将猪的指针交给上层接口.
	//上层接口也可以转换为下层实现类的指针(类型断言)
	p2 := a.(*pig)                    //将接口断言转换为猪类(结构体)
	fmt.Printf("p1=%p p2=%p", p1, p2) //结果:p1=0x598c18 p2=0x598c18   内存地址完全相同.
}

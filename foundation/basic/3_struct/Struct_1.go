package main

import (
	"fmt"
)

/*
Go 语言中的类型可以被实例化，###使用new或&构造的类型实例的类型是"类型的指针"。

结构体成员是由一系列的成员变量构成，这些成员变量也被称为“字段”。字段有以下特性：
字段拥有自己的类型和值。
字段名必须唯一。
字段的类型也可以是结构体，甚至是字段所在结构体的类型。

type 类型名 3_struct{}可以理解为将 3_struct{} 结构体定义为类型名的类型。

结构体的定义只是一种内存布局的描述，只有当结构体实例化时，才会真正地分配内存，因此必须在定义结构体并实例化后才能使用结构体的字段。

实例化就是根据结构体定义的格式创建一份与格式一致的内存区域，结构体实例与实例间的内存是完全独立的。

Go语言可以通过多种方式实例化结构体，根据实际需要可以选用不同的写法。
*/
func main2() {
	//初始化结构体的三种方式.
	initiateStruct()

	//用大括号和键值对初始化结构体的同时赋值.
	assignment()

	//匿名结构体
	anonymousStruct()

	//模拟结构体的构造函数(Go没构函)
	callConstructFunction()

	//用struct {字段1,字段2,..}{初始化1,初始化2,..}的方式,来定义和初始化匿名结构体.
	GenJsonData()
}

/*
*
模拟构造函数.
Go语言中的结构体没有构造函数.
注意: 返回值是指针.无论内部用&还是用new来实例化结构体.

结果:

	构造函数返回的类型是:*main.People,字段name的值是:李四
*/
func callConstructFunction() {
	people := newPeoplePointer("李四")
	fmt.Printf("构造函数返回的类型是:%T,字段name的值是:%s\n", people, people.name)
}

func newPeoplePointer(name string) *People {
	return &People{ //返回指针. 注意,这里即使是用new来初始化,也返回指针类型.
		name: name,
	}
}

/*
*
匿名结构体, 初始化和赋值.

匿名结构体的初始化写法由结构体定义和键值对初始化两部分组成，
结构体定义时没有结构体类型名，只有字段和类型定义，
键值对初始化部分由可选的多个键值对组成，
如下格式所示：

	ins := 3_struct {
		// 匿名结构体字段定义
		字段1 字段类型1
		字段2 字段类型2
		…
	}{
		// 字段值初始化
		初始化字段1: 字段1的值,
		初始化字段2: 字段2的值,
		…
	}

注意:

	键值对初始化部分是可选的，不初始化成员时，匿名结构体的格式变为：
	ins := 3_struct {
		字段1 字段类型1
		字段2 字段类型2
		…
	}

结果:

	传入的匿名结构体的类型是:*3_struct { id int; name string }, 内部的字段是,id:11,name:张三
*/
func anonymousStruct() {
	//实例化一个匿名结构体, 然后传入下面的函数.
	student := struct { //实例化匿名结构体时,不需要写自定义的名字.
		id   int
		name string
	}{
		11,
		"张三",
	}

	//调用下面的函数,以上面初始化的匿名结构体为参数. 下面的函数以匿名结构体为形参.
	withParamOfAnonymousStruct(&student)
}

// 定义一个函数, 参数为一个"匿名结构体"
func withParamOfAnonymousStruct(student *struct { //注意,参数是现定义的一个匿名结构体,占用多行
	id   int
	name string
}) {
	//函数内部只是打印传入的参数而已
	fmt.Printf("传入的匿名结构体的类型是:%T, 内部的字段是,id:%d,name:%s\n", student, student.id, student.name)
}

type People struct {
	name string
	//类型是, 本类型的指针.
	child *People
}

func assignment() {
	relation := &People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}

	fmt.Println("用大括号键值对赋值结构体:", relation)
}

/*
初始化结构体的三种方式.
*/
func initiateStruct() {
	//结构体实例化的方式一: 基本的实例化形式.用: var name T 即可.  T是自定义的结构体名.
	basicInstantiate()

	//结构体实例化的方式二:用new创建指针类型的结构体.
	pointerInstantiate()

	//结构体实例化的方式三:用&取地址的方式
	addressInitiate()
}

/*
*
结构体实例化的方式一: 基本的实例化形式.用: var name T 即可.  T是自定义的结构体名.

结果:

	基本的实例化形式,类型:main.simple1,字段id:1,字段name:name
*/
func basicInstantiate() {
	//用: var name T 即可.  T是自定义的结构体名.
	var a simple1
	//为实例化后的结构体赋值.
	a.id = 1
	a.name = "name"

	fmt.Printf("结构体实例化的方式一: 基本的实例化形式,类型:%T,字段id:%d,字段name:%s", a, a.id, a.name)
	fmt.Println()
}

func pointerInstantiate() {
	/*
	   结构体实例化的方式二:用new创建指针类型的结构体.

	   还可以使用 new 关键字对类型（包括结构体、整型、浮点数、字符串等）进行实例化，结构体在实例化后会形成指针类型的结构体。
	   使用 new 的格式如下：
	   ins := new(T)

	   其中：
	   T 为类型，可以是结构体、整型、字符串等。
	   ins：T 类型被实例化后保存到 ins 变量中，ins 的类型为 *T，属于指针。

	   结果:
	   	结构体实例化的方式二:用new创建指针类型的结构体.,类型:*main.simple1,字段id:1,字段name:name

	*/
	//thePointer变量的类型为 *T，属于指针
	thePointer := new(simple1)

	thePointer.id = 1
	thePointer.name = "name"

	fmt.Printf("结构体实例化的方式二:用new创建指针类型的结构体.,类型:%T,字段id:%d,字段name:%s", thePointer, thePointer.id, thePointer.name)
	fmt.Println()
}

/*
结构体实例化的方式三:用&取地址的方式, 创建指针类型的结构体.
该方式, 等同于用new的方式.
格式:

	ins := &T{}  //记得有大括号

结果:

	结构体实例化的方式三:用&取地址的方式, 创建指针类型的结构体,类型:main.simple1,字段id:1,字段name:name, 字段pointer:1

注意:

###与new不同,返回的变量不是指针.
*/
func addressInitiate() {
	thePointer := simple1{}

	thePointer.id = 1
	thePointer.name = "name"
	var a = 1
	thePointer.pointer = &a

	fmt.Printf("结构体实例化的方式三:用&取地址的方式, 创建指针类型的结构体,类型:%T,字段id:%d,字段name:%s, 字段pointer:%d", thePointer, thePointer.id, thePointer.name, *thePointer.pointer)
	fmt.Println()
}

type simple1 struct {
	id      int
	name    string
	pointer *int
}

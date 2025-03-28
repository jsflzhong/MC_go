package __package

/*
在Go语言中封装就是把抽象出来的字段和对字段的操作封装在一起，数据被保护在内部，程序的其它包只能通过被授权的方法，才能对字段进行操作。

封装的好处：
隐藏实现细节；
可以对数据进行验证，保证数据安全合理。

如何体现封装：
对结构体中的属性进行封装；
通过方法，包，实现封装。

封装的实现步骤：
将结构体、字段的首字母小写；
给结构体所在的包提供一个工厂模式的函数，首字母大写，类似一个构造函数；
提供一个首字母大写的 Set 方法（类似其它语言的 public），用于对属性判断并赋值；
提供一个首字母大写的 Get 方法（类似其它语言的 public），用于获取属性的值。

*/

/*********************************************
注意:
	由于本go文件不在main包下, 本身的package声明也不是main包, 所以下面的函数会在main包那边被使用.
*/

type people struct {
	age  int //其它包不能直接访问.
	name string
}

/*
模拟一个工厂形式的构造函数
*/
func NewPeople() *people {
	return &people{
		age:  0,
		name: "",
	}
}

/*
模拟一个Getter
用"接收器"绑定了上面的type. "接收器"的概念,会在后面讲到.
注意: 首字母必须大写, 否则无法被其他包使用.

[MC]
	为什么要用 *people 而不是 people？
	Go 允许方法绑定到 值类型 和 指针类型，两者区别如下：

	（1）绑定到值类型
	func (p people) GetAge() int {
		return p.age
	}
	p 是 people 类型的一个拷贝，不会影响原来的对象。
	如果 p.age 在方法里被修改，外部的 people 实例不会变。

	（2）绑定到指针类型
	func (p *people) SetAge(age int) {
		p.age = age
	}
	p 是 people 的 指针，它直接指向原对象。
	方法里修改 p.age，外部 people 的 age 也会改变。
	*/
func (p *people) GetAge() int {
	//注意,由于字段首小,所以只能在本包中这样调用.
	return p.age
}

/*
模拟一个Setter
用接收器绑定了上面的type.
注意: 首字母必须大写, 否则无法被其他包使用.
*/
func (p *people) SetAge(a int) {
	p.age = a
}

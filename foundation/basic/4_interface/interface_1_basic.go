package main

/*
1.每个接口类型由数个方法组成。接口的形式代码如下：

	Go语言的接口在命名时，一般会在单词后面添加 er，如有写操作的接口叫 Writer，有字符串功能的接口叫 Stringer，
	有关闭功能的接口叫 Closer 等。

	type 接口类型名 interface{
		//当方法名首字母是大写时，且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
		方法名1( 参数列表1 ) 返回值列表1
		//参数列表和返回值列表中的参数变量名可以被忽略, 例如:Write([]byte) error
		方法名2( 参数列表2 ) 返回值列表2
		…
	}

2.Go语言提供的很多包中都有接口，例如 io 包中提供的 Writer 接口：

	type Writer interface {
		Write(p []byte) (n int, err error)
	}
	这个接口可以调用 Write() 方法写入一个字节数组（[]byte），返回值告知写入字节数（n int）和可能发生的错误（err error）。

3.类似的，还有将一个对象以字符串形式展现的接口，只要实现了这个接口的类型，
在调用 String() 方法时，都可以获得对象对应的字符串。在 fmt 包中定义如下：

	type Stringer interface {
		String() string
	}
	Stringer 接口在Go语言中的使用频率非常高，功能类似于 Java 或者 C# 语言里的 ToString 的操作。

*/
func main1() {

}

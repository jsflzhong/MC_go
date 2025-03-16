package main

import (
	"encoding/json"
	"fmt"
)

/*
*
结构体和json之间的相互转换.
内置API:

	json->3_struct:
		jsonData, _ := json.Marshal(struct_)
	3_struct->json:
		json.Unmarshal(jsonData, &struct_screenAndTouch)

注意:

	反序列化可以只反序列化一部分,不用全字段.

###注意:

	如果想让json.Marshal(Pojo)生效, 则结构体中的字段名要首大.

结果:

	@@@structToJason(): {"Size":5.5,"ResX":1920,"ResY":1080,"Capacity":2910,"HasTouchID":true}
	@@@jsonToStruct():{Screen:{Size:5.5 ResX:1920 ResY:1080} HasTouchID:true}
	@@@jsonToStruct2:{Battery:{Capacity:2910} HasTouchID:true}
*/
func main5() {
	//把结构体序列化为JSON
	jsonData := structToJason()

	//把JSON反序列化为结构体(只取部分字段来反序列化)
	jsonToStruct(jsonData)

	//把JSON反序列化为结构体(只取另一部分字段来反序列化)
	jsonToStruct2(jsonData)

	//使用关键字"json:xx"来改变 结构体->JSON 时的JSON中的key的名字.注意,要想转json, 字段名必须首大!
	changeJsonKeyName()

	//在字段后加入 omitempty（使用逗号,与前面的内容分隔），来过滤掉转换的 JSON 格式中的空值
	filterOutNilValues()
}

/*
在字段后加入 omitempty（使用逗号,与前面的内容分隔），来过滤掉转换的 JSON 格式中的空值

运行结果:

	{
	    "Name":"cow boy",
	    "Age":37,
	    "Skills":[
	        {
	            "level":1	//注意,这里的空字段被忽略转换了.
	        },
	        {
	            "name":"Flash your dog eye",
	            "level":0
	        },
	        {
	            "name":"Time to have Lunch",
	            "level":3
	        }
	    ]
	}
*/
func filterOutNilValues() {
	// 声明技能结构体
	type Skill struct {
		//如果想让json.Marshal(Pojo)生效, 则结构体中的字段名要首大.
		Name  string `json:"name,omitempty"` //关键字omitempty
		Level int    `json:"level"`
	}
	// 声明角色结构体
	type Actor struct {
		//如果想让json.Marshal(Pojo)生效, 则结构体中的字段名要首大.
		Name   string
		Age    int
		Skills []Skill
	}
	// 填充基本角色数据
	a := Actor{
		Name: "cow boy",
		Age:  37,
		Skills: []Skill{
			{Name: "", Level: 1},
			{Name: "Flash your dog eye"},
			{Name: "Time to have Lunch", Level: 3},
		},
	}
	result, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	jsonStringData := string(result)
	fmt.Println(jsonStringData)
}

/*
使用关键字"json:xx"来改变 结构体->JSON 时的JSON中的key的名字.

结果:

	{
	    "Name":"cow boy",
	    "Age":37,
	    "Skills":[
	        {
	            "name":"Roll and roll",
	            "level":1
	        },
	        {
	            "name":"Flash your dog eye",
	            "level":2
	        },
	        {
	            "name":"Time to have Lunch",
	            "level":3
	        }
	    ]
	}
*/
func changeJsonKeyName() {
	// 声明技能结构体
	type Skill struct {
		Name  string `json:"name"` //用关键字json:, 来改变JSON序列化后的key的名字.注意,要想转json, 字段名必须首大!
		Level int    `json:"level"`
	}
	// 声明角色结构体
	type Actor struct {
		Name   string
		Age    int
		Skills []Skill
	}
	// 填充基本角色数据
	a := Actor{
		Name: "cow boy",
		Age:  37,
		Skills: []Skill{
			{Name: "Roll and roll", Level: 1},
			{Name: "Flash your dog eye", Level: 2},
			{Name: "Time to have Lunch", Level: 3},
		},
	}
	//转换
	result, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	//string化
	jsonStringData := string(result)
	fmt.Println("@@@changeJsonKeyName:", jsonStringData)
}

/*
定义结构体: 手机屏幕
*/
type Screen struct {
	Size       float32 // 屏幕尺寸
	ResX, ResY int     // 屏幕水平和垂直分辨率
}

/*
定义结构体: 电池
*/
type Battery struct {
	Capacity int // 容量
}

/*
定义函数: 把结构体->转换生成json数据
*/
func GenJsonData() []byte {
	// 定义并初始化一个匿名结构体
	struct_ := &struct {
		Screen
		Battery
		HasTouchID bool // 序列化时添加的字段：是否有指纹识别
	}{ //开始初始化
		// 屏幕参数
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		// 电池参数
		Battery: Battery{
			2910,
		},
		// 是否有指纹识别
		HasTouchID: true,
	}

	// 调用内置API, 将数据序列化为json
	jsonData, _ := json.Marshal(struct_)
	return jsonData
}

/*
*
把JSON反序列化为结构体(只取另一部分字段来反序列化)
*/
func jsonToStruct2(jsonData []byte) {
	// 只需要电池和指纹识别信息的结构和实例
	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}
	// 反序列化到batteryAndTouch
	json.Unmarshal(jsonData, &batteryAndTouch)
	// 输出screenAndTouch的详细结构
	fmt.Printf("@@@jsonToStruct2:%+v\n", batteryAndTouch)
}

/*
*
把JSON反序列化为结构体(只取部分字段来反序列化)
*/
func jsonToStruct(jsonData []byte) {
	// 只需要屏幕和指纹识别信息的结构和实例
	struct_screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}
	// 反序列化到screenAndTouch. 注意,结构体中没定义"Battery"类型的字段 即反序列化时可以只反序列化一部分.
	json.Unmarshal(jsonData, &struct_screenAndTouch)
	// 输出screenAndTouch的详细结构
	fmt.Printf("@@@jsonToStruct():%+v\n", struct_screenAndTouch)
}

/*
*
把结构体序列化为JSON
*/
func structToJason() []byte {
	// 生成一段json数据
	jsonData := GenJsonData()
	fmt.Println("@@@structToJason():", string(jsonData))
	return jsonData
}

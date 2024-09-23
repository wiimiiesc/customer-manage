package view

import (
	"customerManage/service"
	"fmt"
)

type CustomerView struct {
	key            int  // 用户输入
	loop           bool // 是否退出
	customerServer service.CustomerService
}

func NewCustomerView() *CustomerView {
	customerServer := service.NewCustomerService()
	return &CustomerView{
		key:            0,
		loop:           false,
		customerServer: *customerServer,
	}
}

func (this CustomerView) MainMune() {

	for {
		fmt.Println("\n\n------------------用户信息管理软件--------------------")
		fmt.Println("                  1 添 加 用 户")
		fmt.Println("                  2 修 改 用 户")
		fmt.Println("                  3 删 除 用 户")
		fmt.Println("                  4 客 户 列 表")
		fmt.Println("                  5 退	出")
		fmt.Print("\n请选择(1-5)：")

		fmt.Scanln(&this.key)

		switch this.key {
		case 1: // 添加用户
			this.add()
		case 2: // 修改用户
			this.change()
		case 3: // 删除用户
			this.delete()
		case 4: // 客户列表
			this.list()
		case 5: // 退出
			this.exit()
		default:
			fmt.Println("输入不规范，请重新输入...")

		}
		this.key = -1
		// 判断是否退出
		if this.loop {
			fmt.Println("已退出程序")
			break
		}
	}

}

// 退出程序
func (this *CustomerView) exit() {
	var choice string
	for {
		fmt.Print("确定要退出吗(Y/N)：")
		fmt.Scanln(&choice)

		if choice == "Y" || choice == "N" || choice == "y" || choice == "n" {
			break
		}
	}
	if choice == "Y" || choice == "y" {
		this.loop = true
	}

}

// 修改用户

func (this *CustomerView) change() {
	// var name string
	// var age int
	// var gender string
	// var phone string
	// var email string
	id := -1

	fmt.Println("\n\n\n--------------------修改用户----------------------")
	fmt.Println("请输入修改用户编号(-1 退出)：")
	fmt.Scanln(&id)

	name, gender, age, phone, email, err := this.customerServer.GetInfoById(id)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("姓名(%v)：", name)
	fmt.Scanln(&name)
	fmt.Printf("年龄(%v)：", age)
	fmt.Scanln(&age)
	fmt.Printf("性别(%v)：", gender)
	fmt.Scanln(&gender)
	fmt.Printf("手机号(%v)：", phone)
	fmt.Scanln(&phone)
	fmt.Printf("邮箱(%v)：", email)
	fmt.Scanln(&email)

	err = this.customerServer.ChangeById(id, name, gender, age, phone, email)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("--------------------修改完成----------------------")

}

// 删除用户
func (this *CustomerView) delete() {
	var num int
	fmt.Println("\n\n\n--------------------删除用户----------------------")
	fmt.Print("请选择待删除用户编号(-1 退出)：")
	fmt.Scanln(&num)
	if num == -1 {
		return
	}

	err := this.customerServer.Delete(num)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("--------------------删除完成----------------------")

}

// 新增用户
func (this *CustomerView) add() {
	var name string
	var age int
	var gender string
	var phone string
	var email string

	fmt.Println("\n\n\n--------------------添加用户----------------------")
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	fmt.Print("手机号：")
	fmt.Scanln(&phone)
	fmt.Print("邮箱：")
	fmt.Scanln(&email)

	this.customerServer.Add(name, gender, age, phone, email)

	fmt.Println("--------------------添加完成----------------------")

}

// 客户列表
func (this *CustomerView) list() {
	fmt.Println("\n\n\n--------------------客户列表----------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t手机号\t\t邮箱")
	list := this.customerServer.List()
	for _, user := range list {
		fmt.Println(user.GetInfo())
	}
	fmt.Println("-------------------客户列表完成-------------------")
}

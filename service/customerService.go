package service

import (
	"customerManage/model"
	"errors"
)

type CustomerService struct {
	// 用于存储用户数据切片
	customers []*model.Customer
	// 记录当前的用户id
	customerNum int
}

// 构造函数
func NewCustomerService() *CustomerService {
	customerService := CustomerService{}
	user := model.NewCustomer(0, "张三", "男", 20, "19525658745", "123@qq.com")
	customerService.customers = append(customerService.customers, user)
	return &customerService
}

// 获取用户信息
func (this *CustomerService) GetInfoById(id int) (name string, gender string, age int, phone string, email string, err error) {
	for _, user := range this.customers {
		if user.Id == id {
			name = user.Name
			gender = user.Gender
			age = user.Age
			phone = user.Phone
			email = user.Email
			err = nil
			return
		}
	}
	err = errors.New("未找到修改用户")
	return
}

// 修改用户
func (this *CustomerService) ChangeById(id int, name string, gender string, age int, phone string, email string) error {
	for _, user := range this.customers {
		if user.Id == id {
			user.ChangeInfo(name, gender, age, phone, email)
			return nil
		}
	}

	return errors.New("未找到用户...")
}

// 删除用户
func (this *CustomerService) Delete(id int) error {
	index := this.FindIndexById(id)
	if index == -1 {
		return errors.New("未找到该用户...")
	}

	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return nil
}

// 查找用户
func (this CustomerService) FindIndexById(id int) int {
	for index, user := range this.customers {
		if user.Id == id {
			return index
		}
	}
	return -1
}

// 返回用户列表
func (this *CustomerService) List() []*model.Customer {
	return this.customers
}

// 增加用户
func (this *CustomerService) Add(name string, gender string, age int, phone string, email string) error {
	this.customerNum += 1
	this.customers = append(this.customers, model.NewCustomer(
		this.customerNum, name, gender, age, phone, email,
	))

	return nil
}

package model

import (
	"fmt"
)

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 工厂构造函数
func NewCustomer(id int, name string, gender string, age int, phone string, email string) *Customer {
	return &Customer{Id: id, Name: name, Age: age, Phone: phone, Email: email, Gender: gender}

}

func (this *Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Gender, this.Age, this.Phone, this.Email)
}

func (this *Customer) ChangeInfo(name string, gender string, age int, phone string, email string) {
	this.Name = name
	this.Gender = gender
	this.Age = age
	this.Phone = phone
	this.Email = email
}

package main

import "fmt"

type User struct {
	ID int
	FirstName string
	Pay *Payment
}

type Payment struct {
	Salary float64
	Bonus float64
}

func NewUser() *User {
	u := new(User)
	u.Pay = new(Payment)
	return u
}

func NewPayment() *Payment {
	p := new(Payment)
	return p
}

func main() {
	user1 := &User{
		ID: 1,
		FirstName: "ömer",
		Pay: &Payment{
			Salary: 500,
			Bonus: 300,
		},
	}
	fmt.Println(user1.Pay)
}

func (u User) GetFullName() string {
	return u.FirstName + " -<>-"
}

func (u User) GetUserName() string {
	return u.FirstName
}
package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) save() {
	fmt.Println("Inside save method", u)
}

func (u user) isLegalAge() bool {
	return u.age > 18
}

func (u *user) makeBirthday() {
	u.age++
}

func main() {
	user := user{"Filipe", 26}
	user.save()

	isLegalAge := user.isLegalAge()
	fmt.Println("isLegalAge", isLegalAge)
	user.makeBirthday()
	fmt.Println(user.age)
}

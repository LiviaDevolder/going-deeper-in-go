package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Printf("Saving user %s in database", u.name)
}

func (u user) ofLegalAge() bool {
	return u.age >= 18
}

func (u *user) makeBirthday() {
	u.age++
}

func main() {
	user1 := user{"User 1", 22}
	fmt.Println(user1)

	user1.save()

	user2 := user{"User 2", 23}
	user2.save()
	user2.ofLegalAge()
	user2.makeBirthday()

	fmt.Println(user2.age)
}

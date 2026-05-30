package main

import "fmt"

type User struct {
	Name  string
	Age   int
	Email string
}

type Address struct {
	City    string
	State   string
	Country string
}

type Employee struct {
	ID      int
	User    User
	Address Address
}

/* Value Receiver
(u User) is called the receiver, think of it as "attach this function to User"
*/
func (u User) Greeting() string {
	return "Hello, my name is " + u.Name
}

// Value Receiver
func (u User) changeName(newName string) {
	u.Name = newName
	fmt.Println(u)
}

// Pointer Receiver
func (u *User) changeNameByRef(newName string) {
	u.Name = newName
	fmt.Println(u.Name)
}

// Pointer Receiver
func (u *User) Birthday() {
	u.Age++
}


func (u User) IsAdult() bool {
	return u.Age >= 18
}

func main() {
	user := User{
		Name:  "Atib",
		Age:   35,
		Email: "atib@example.com",
	}

	fmt.Println(user)

	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)

	user.Age = 36

	fmt.Println("Updated Age:", user.Age)

	employee := Employee{
		ID:   1001,
		User: user,
		Address: Address{
			City:    "Bangalore",
			State:   "Karnataka",
			Country: "India",
		},
	}
	fmt.Println(employee)

	var emp_x Employee = Employee{
		ID: 1002,
		User: User{
			Name: "X",
			Email: "X@gmail.com",
			Age: 40,
		},
		Address: Address{
			City: "Chennai",
			Country: "India",
			State: "Tamil Nadu",
		},
	}
	fmt.Println(emp_x)

	// You can now do user.Greeting() rather than Greeting(user)
	fmt.Println(user.Greeting())

	fmt.Println(user)
	user.changeName("John")
	// This doesn't reflect the change in name since the VALUE was copied
	fmt.Println(user)

	// This reflects the change since the pointer reference was passed
	user.changeNameByRef("John")
	fmt.Println(user)

	// Another example of pass by reference
	fmt.Println("Age:", user.Age)
	user.Birthday()
	fmt.Println("Age:", user.Age)

	fmt.Println(user.IsAdult())
}
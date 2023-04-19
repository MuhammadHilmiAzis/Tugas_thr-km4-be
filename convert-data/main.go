package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func parseUser(data string) User {
	userData := strings.Split(data, ",")
	age, _ := strconv.Atoi(userData[1])
	return User{Name: userData[0], Age: age, Address: userData[2]}
}

func main() {
	data := "Edo,20,Jakarta"
	user := parseUser(data)
	fmt.Printf("%+v\n", user) // Output: {Name: "Edo", Age: 20, Address: "Jakarta"}

	data = "Budi,30,Bandung"
	user = parseUser(data)
	fmt.Printf("%+v\n", user) // Output: {Name: "Budi", Age: 30, Address: "Bandung"}
}

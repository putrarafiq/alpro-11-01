package main

import "fmt"

func main() {
	var name string

	name = "Rafiq Khairan Putra Permana"

	fmt.Println("Nama : ", name)

	var lastName = "Permana"
	fmt.Println("Nama Belakang : ", lastName)

	middleName := "Khairan"
	fmt.Println("Nama Tengah : ", middleName)

	var(
		fullName = "Rafiq Khairan Putra Permana"
		firstName = "Rafiq"
	)

	fmt.Println(fullName)
	fmt.Println(firstName)
}

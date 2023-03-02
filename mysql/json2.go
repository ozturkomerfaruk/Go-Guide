package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	ID      int
	Kind    string
	Address string
}

type Interest struct {
	ID   int
	Name string
}

type Person struct {
	ID        int
	Firstname string
	Lastname  string
	Username  string
	Gender    string
	Name      Name
	Email     []Email
	Interest  []Interest
}

func main() {
	person := Person{
		ID:        1,
		Firstname: "Ömer Faruk",
		Lastname:  "Öztürk",
		Username:  "ozturkomerfaruk",
		Gender:    "Erkek",
		Name: Name{
			Family:   "Öztürk",
			Personal: "Muzaffer",
		},
		Email: []Email{
			{ID: 1, Kind: "Work", Address: "iletisim@omerfarukozturk.com"},
			{ID: 2, Kind: "Personal", Address: "omerfarukozturk026@gmail.com"},
		},
		Interest: []Interest{
			{ID: 1, Name: "Müzik"},
			{ID: 2, Name: "Kitap"},
		},
	}

	WriteMessage("Reading Operation Started")
	WriteMessage("Personal Fullname")
	WriteStarLine()

	res := GetPerson(&person)
	WriteMessage(res)
	WriteStarLine()

	WriteMessage("Personal Email with Index")
	WriteStarLine()
	resEmail := GetPersonEmailAddress(&person, 0)
	WriteMessage(resEmail)
	WriteStarLine()

	WriteMessage("Personal Email Object With Index")
	WriteStarLine()

	resEmail2 := GetPersonEmail(&person, 1)
	WriteMessage(resEmail2.Address)
	WriteStarLine()

	WriteMessage("Reading Operation Ended")

	WriteMessage("Writing Operation Started")
	SaveJSON("person.json", person)
	WriteMessage("Writing Operation Ended")
}

func GetPerson(p *Person) string {
	return p.Firstname + " " + p.Lastname
}

func GetPersonEmailAddress(p *Person, i int) string {
	return p.Email[i].Address
}

func GetPersonEmail(p *Person, i int) Email {
	return p.Email[i]
}

func WriteMessage(msg string) {
	fmt.Println(msg)
}

func WriteStarLine() {
	fmt.Println("* * * * * * * * * * * * * * * * * *")
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal Error : ", err.Error())
		os.Exit(1)
	}
}

func SaveJSON(filename string, key interface{}) {
	outfile, err := os.Create(filename)
	CheckError(err)
	encoder := json.NewEncoder(outfile)
	err = encoder.Encode(key)
	CheckError(err)
	outfile.Close()
}

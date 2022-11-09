package main

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/jaswdr/faker"
	"github.com/urfave/cli/v2"
)

func printHeading(text string) {
	color.New(color.Italic, color.Bold, color.FgMagenta).Printf("--------%s--------\n", text)
}

func printItem(title string, text string) {
	color.New(color.Bold).Printf("%s: \t", title)
	fmt.Println(text)
}

func generate(ctx *cli.Context) error {
	fake := faker.New()

	firstName := fake.Person().FirstName()
	lastName := fake.Person().LastName()
	email := fake.Person().Contact().Email
	phone := fake.Person().Contact().Phone
	password := fake.Internet().Password()
	country := fake.Address().Country()
	city := fake.Address().City()
	street := fake.Address().StreetAddress()
	zip := randomZip()

	printHeading("General information")
	printItem("First name", firstName)
	printItem("Last name", lastName)
	printItem("Email", "\t"+email)
	printItem("Password", password)
	printItem("Phone", "\t"+phone)
	printItem("Country", country)
	printItem("City", "\t"+city)
	printItem("Street", street)
	printItem("ZIP code", strconv.Itoa(zip))

	return nil
}

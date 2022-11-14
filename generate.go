package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/jaswdr/faker"
	"github.com/urfave/cli/v2"
)

func printHeading(text string) {
	color.New(color.Italic, color.Bold, color.FgMagenta).Printf("%s\n", text)
}

func printItem(title string, text string) {
	color.New(color.Bold).Printf("%s: \t", title)
	fmt.Println(text)
}

func generate(ctx *cli.Context) error {
	rand.Seed(time.Now().UnixNano())
	fake := faker.New()

	firstName := fake.Person().FirstName()
	lastName := fake.Person().LastName()
	email := randomEmail(&firstName)
	phone := fake.Person().Contact().Phone
	password := fake.Internet().Password()
	country := fake.Address().Country()
	city := fake.Address().City()
	street := fake.Address().StreetAddress()
	zip := randomZip()
	cardNumber := randomCardNumber(mastercardPrefixArr)
	endDate := faker.New().Payment().CreditCardExpirationDateString()

	printHeading("--------General information--------")
	printItem("First name", firstName)
	printItem("Last name", lastName)
	printItem("Email", "\t"+email)
	printItem("Password", password)
	printItem("Phone", "\t"+phone)
	printHeading("------------Address----------------")
	printItem("Country", country)
	printItem("City", "\t"+city)
	printItem("Street", strings.TrimPrefix(street, "%"))
	printItem("ZIP code", strconv.Itoa(zip))
	printHeading("--------Credit/Debit card----------")
	printItem("Card number", cardNumber)
	printItem("End date", endDate)
	printItem("CVV", "\t"+endDate)
	printItem("Card type", endDate)

	return nil
}

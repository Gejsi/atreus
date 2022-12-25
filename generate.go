package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/jaswdr/faker"
	"github.com/jwalton/gchalk"
	"github.com/urfave/cli/v2"
)

func printHeading(text string) {
	fmt.Println(gchalk.WithItalic().WithMagenta().Bold(text))
}

func printItem(title string, text string) {
	fmt.Printf(gchalk.Bold("%s: \t"), title)
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
	cardNumber, cardIssuer := randomCardNumber()
	endDate := faker.New().Payment().CreditCardExpirationDateString()
	cvv := randomCVV()

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
	printHeading("-----------Credit card-------------")
	printItem("Card number", cardNumber)
	printItem("End date", endDate)
	printItem("CVV", "\t"+strconv.Itoa(cvv))
	printItem("Card type", cardIssuer)

	return nil
}

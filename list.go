package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/jwalton/gchalk"
	"github.com/urfave/cli/v2"
)

const baseUrl string = "https://www.1secmail.com/api/v1/?action="

type Email struct {
	Id      int    `json:"id"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Date    string `json:"date"`
}

func list(ctx *cli.Context) error {
	emailArg, err := mail.ParseAddress(ctx.String("email"))
	if err != nil {
		return cli.Exit("Please, provide a valid email.", 1)
	}

	splitEmail := strings.Split(emailArg.Address, "@")
	name := splitEmail[0]
	domain := splitEmail[1]

	url := baseUrl + "getMessages&login=" + name + "&domain=" + domain

	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var emails []Email
	if err := json.Unmarshal(body, &emails); err != nil {
		log.Fatalln("Cannot unmarshal JSON", err)
	}

	for _, email := range emails {
		fmt.Printf(gchalk.WithGreen().Bold("%d"), email.Id)
		fmt.Printf(" ")
		fmt.Printf(gchalk.WithYellow().Underline("%s"), email.From)
		fmt.Printf(" ")
		fmt.Printf(gchalk.BrightBlack("%s"), email.Date)
		fmt.Printf(" ")
		fmt.Printf("\"%s\"", email.Subject)
		fmt.Println()
	}

	return nil
}

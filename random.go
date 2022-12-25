package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func randomInRange(min int, max int) int {
	return rand.Intn(max-min) + min
}

func randomZip() int {
	return randomInRange(10000, 99999)
}

func randomEmail(firstName *string) string {
	domains := []string{"com", "net", "org"}
	name := *firstName + strconv.Itoa(rand.Intn(300))

	return strings.ToLower(name) + "@1secmail." + domains[rand.Intn(len(domains))]
}

var visaPrefixArr = &[]string{
	"4539",
	"4556",
	"4916",
	"4532",
	"4929",
	"40240071",
	"4485",
	"4716",
	"4",
}

var mastercardPrefixArr = &[]string{
	"51",
	"52",
	"53",
	"54",
	"55",
}

var discoverPrefixArr = &[]string{"6011"}

var jcbPrefixArr = &[]string{"35"}

var cardIssuers = &[4]string{"Visa", "MasterCard", "Discover", "JCB"}

func reverse(arr *[]int) []int {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}

	return *arr
}

func calculateCardNumber(prefix *string) string {
	cardNumber := make([]int, len(*prefix))
	const length int = 16

	for i, char := range *prefix {
		cardNumber[i] = int(char - '0')
	}

	for len(cardNumber) < (length - 1) {
		cardNumber = append(cardNumber, rand.Intn(10))
	}

	cardNumber = reverse(&cardNumber)

	var sum int

	for i := 0; i < length-1; i += 2 {
		odd := cardNumber[i] * 2

		if odd > 9 {
			odd -= 9
		}

		sum += odd

		if i != length-2 {
			sum += cardNumber[i+1]
		}
	}

	checkDigit := ((int(math.Floor(float64(sum)/10))+1)*10 - sum) % 10
	cardNumber = reverse(&cardNumber)
	cardNumber = append(cardNumber, checkDigit)

	// one-liner that converts []int to string (without a delimiter)
	res := strings.Trim(strings.Join(strings.Split(fmt.Sprint(cardNumber), " "), ""), "[]")

	return res
}

func randomCardNumber() (string, string) {
	issuer := cardIssuers[rand.Intn(len(cardIssuers))]
	var cardNumber string

	if issuer == "MasterCard" {
		cardNumber = calculateCardNumber(&(*mastercardPrefixArr)[rand.Intn(len((*mastercardPrefixArr)))])
	} else if issuer == "Discover" {
		cardNumber = calculateCardNumber(&(*discoverPrefixArr)[rand.Intn(len((*discoverPrefixArr)))])
	} else if issuer == "JCB" {
		cardNumber = calculateCardNumber(&(*jcbPrefixArr)[rand.Intn(len((*jcbPrefixArr)))])
	} else {
		cardNumber = calculateCardNumber(&(*visaPrefixArr)[rand.Intn(len((*visaPrefixArr)))])
	}

	return cardNumber, issuer
}

func randomCVV() int {
	return randomInRange(100, 999)
}

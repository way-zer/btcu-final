package main

import (
	"btcu-final/client"
	"log"
)

func main() {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	println("GenerateKeys")
	pri, pub, err := client.GenerateKeys()
	check(err)

	println("Register")
	const TEST_HASH = "123456789123456789123456798"
	err = client.Register(&client.Copyright{
		Name:      "TEST",
		Author:    "TEST_USER",
		Press:     "TEST_PRESS",
		Hash:      TEST_HASH,
		PublicKey: *pub,
	}, *pri)
	check(err)

	println("GetRightByHash")
	data, err := client.GetRightByHash(TEST_HASH)
	check(err)
	println(data)

	println("GetRightByInfo")
	data, err = client.GetRightByInfo("TEST", "TEST_USER", "TEST_PRESS")
	check(err)
	println(data)

	println("===TEST PASS===")
}

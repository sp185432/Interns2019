package main

import (
	"testing"
)

var testcases = []struct {
	input, passphrase, output string
}{
	{input: "SHARANYA", passphrase: "password", output: "SHARANYA"},
	{input: "destination", passphrase: "password", output: "destination"},
	//{input: "destination", passphrase: "key", output: "destination"},
	{input: "journey8818", passphrase: "password", output: "journey8818"},
	//{input: "journey8818", passphrase: "password", output: "JOURNEY8818"},
}

//----BASIC TESTING FOR ENCRYPTION AND DECRYPTION--------
func TestEncrypt(t *testing.T) {

	for _, tc := range testcases {

		actual := CreateHash(tc.passphrase)

		expected := "5f4dcc3b5aa765d61d8327deb882cf99"

		if actual != expected {

			t.Fatal("Different Hash value is achieved ")
			t.Fatalf("Actual: %s Expected: %s", actual, expected)

		} else {

			t.Log("The Hash value is obained as expected!")
			t.Logf("Actual: %s Expected: %s", actual, expected)

		}

		ciphertext := Encrypt([]byte(tc.input), tc.passphrase)

		actual1 := Decrypt(ciphertext, tc.passphrase)

		expected1 := tc.output

		if string(actual1) != expected1 {

			t.Fatal("Encryption and Decryption is not achieved ")

		} else {

			t.Log("Encryption and Decryption are obained as expected!")

		}
	}
}

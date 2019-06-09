package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

//CreateHash function takes a key of string type to obtain a hash value
func CreateHash(key string) string {

	//create a new hash
	hasher := md5.New()

	//adding data, write expects bytes, as we have a string use []byte to coerce it to bytes
	hasher.Write([]byte(key))

	//fmt.Println("hash value is: ", hex.EncodeToString(hasher.Sum(nil)))
	//extracting the checksum by calling its sum function and returning hexadecimal value
	return hex.EncodeToString(hasher.Sum(nil))
}

//Encrypt FUNCTION TO ENCRYPT THE GIVEN INPUT
func Encrypt(data []byte, passphrase string) []byte {

	fmt.Println("------ENCRYPTION PROGRAM-------")

	//creating a new block cipher based on the hashed passphrase
	block, _ := aes.NewCipher([]byte(CreateHash(passphrase)))

	/*we have our block cipher, wrap it in a GCM(Galois counter mode), this is a mode of
	operation for symmetric key cryptographic block ciphers*/
	gcm, _ := cipher.NewGCM(block)

	//creating a new byte array the size of nonce which must be passed to seal
	nonce := make([]byte, gcm.NonceSize())

	//populates our nonce with a cryptographically secure random sequence
	io.ReadFull(rand.Reader, nonce)

	/*here we encrypt our text making use of seal function. seal encrypts and authenticates
	plaintext, authenticates the additional data and appends the result to the destination
	returning the updated slice. The nonce must be nonceSize() bytes long and
	should be unique all the time for a given keys */
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	fmt.Println("Encrypted string is: ", string(ciphertext))
	fmt.Println("Length of the encrypted string is: ", len(ciphertext))

	return ciphertext

}

//Decrypt function for performing decryption
func Decrypt(data []byte, passphrase string) []byte {

	fmt.Println("------DECRYPTION PROGRAM-------")

	//creating a hash
	key := []byte(CreateHash(passphrase))

	//creating a new block cipher based on hashed passphrase
	block, _ := aes.NewCipher(key)

	//we have our block cipher, wrap it in a GCM(Galois counter mode)
	gcm, _ := cipher.NewGCM(block)

	//as we wrapped it in gcm mode, obtaining the nonce size
	nonceSize := gcm.NonceSize()

	//separating the nonce and the encrypted data
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	//decrypting and returning it a plaintext
	//Open decrypts and authenticates ciphertext, authenticates the
	// additional data and, if successful, appends the resulting plaintext
	// to dst, returning the updated slice. The nonce must be NonceSize()
	// bytes long and both it and the additional data must match the
	// value passed to Seal.
	// To reuse ciphertext's storage for the decrypted output, use ciphertext[:0]
	// as dst.
	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)

	fmt.Println("Decrypted string is : ", string(plaintext))
	fmt.Println("Length of the decrypted string is:", len(plaintext))

	return plaintext
}

/*func main() {
	var c int
	var input string

	fmt.Println("Enter the input string: ")
	fmt.Scan(&input)

	//consoleReader := bufio.NewReader(os.Stdin)
	//input, _ := consoleReader.ReadString('\n')
	fmt.Println("Entered input string is: ", input)

	//calling the encrypt function
	ciphertext := Encrypt([]byte(input), "password")

	//calling the decrypt function
	plaintext := Decrypt(ciphertext, "password")

	fmt.Println("------COMPARISON OF USER INPUT AND DECRYPTED OUTPUT-------")
	//comparing the input given by user and output obtained by decrypt function
	c = strings.Compare(input, string(plaintext))
	if c == 0 {
		fmt.Println("strings are equal")
	} else {
		fmt.Println("strings are not equal")
	}
}
*/

package main

import (
	"os"

	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
)

func main() {
	sshConfig, err := auth.PrivateKey("root", "/Users/vagrant/.ssh/authorized_keys", ssh.InsecureIgnoreHostKey())
	checkError(err)

	scpClient := scp.NewClient("10.0.2.15:22", &sshConfig) //[remore-server-ip]

	err = scpClient.Connect()
	checkError(err)

	fileData, err := os.Open("/Users/Navya/go/src/files/readfile/file1.txt")
	checkError(err)

	scpClient.CopyFile(fileData, "/var/www/html/test/test.txt", "0655")

	defer scpClient.Session.Close()
	defer fileData.Close()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
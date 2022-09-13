package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const KEY = "e7509a8c032f3bc2a8df1df476f8ef03436185fa"

func EncryptDecrypt(input []byte, key string) (output []byte) {
	out := make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		out[i] = input[i] ^ key[i%len(key)]
	}

	return out
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: phantom-decoder <capture_path>")
	}

	filepath := os.Args[1]

	capData, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("unable to open handle to capture: %w", err)
	}

	plain := EncryptDecrypt(capData, KEY)

	plainFile, err := os.OpenFile("decrypted.png", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatal("unable to open file to write decrypted payload", err)
	}

	nBytes, err := plainFile.Write(plain)
	if err != nil {
		log.Fatal("unable to write decrypted payload", err)
	}

	fmt.Printf("Wrote %d bytes to decrypted.png", nBytes)
}

package main

import (
	"fmt"
	"go-tamboon/cipher"
	"io"
	"os"
)

type Donation struct {
	Name           string `csv:"Name"`
	AmountSubunits int    `csv:"AmountSubunits"`
	CCNumber       string `csv:"CCNumber"`
	CVV            string `csv:"CVV"`
	ExpMonth       int    `csv:"ExpMonth"`
	ExpYear        int    `csv:"ExpYear"`
}

func main() {
	cwd, _ := os.Getwd()
	fmt.Println("Current working directory:", cwd)

	encryptedFile := "data/fng.1000.csv.rot128"
	file, err := os.Open(encryptedFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	outputFile := "data/fng.csv"
	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer out.Close()

	reader, err := cipher.NewRot128Reader(file)

	_, err = io.Copy(out, reader)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return
	}

	fmt.Println("Decryption successful! Output file:", outputFile)
}

package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	encodeCmd := flag.NewFlagSet("encode", flag.ExitOnError)
	decodeCmd := flag.NewFlagSet("decode", flag.ExitOnError)

	inputFile := flag.String("i", "", "Input file")
	outputFile := flag.String("o", "", "Output file")

	if len(os.Args) < 2 {
		fmt.Println("Usage: encoder64 [encode|decode] -i <inputfile> -o <outputfile>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "encode":
		encodeCmd.StringVar(inputFile, "i", "", "Input file")
		encodeCmd.StringVar(outputFile, "o", "", "Output file")
		err := encodeCmd.Parse(os.Args[2:])
		if err != nil {
			return
		}
		if *inputFile == "" {
			fmt.Println("Input file is required for encoding.")
			os.Exit(1)
		}
		encodeFile(*inputFile, *outputFile)

	case "decode":
		decodeCmd.StringVar(inputFile, "i", "", "Input file")
		decodeCmd.StringVar(outputFile, "o", "", "Output file")
		err := decodeCmd.Parse(os.Args[2:])
		if err != nil {
			return
		}
		if *inputFile == "" {
			fmt.Println("Input file is required for decoding.")
			os.Exit(1)
		}
		decodeFile(*inputFile, *outputFile)

	default:
		fmt.Println("Usage: encoder64 [encode|decode] -i <inputfile> -o <outputfile>")
		os.Exit(1)
	}
}

func encodeFile(inputFile, outputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	encodedData := base64.StdEncoding.EncodeToString(data)

	if outputFile == "" {
		outputFile = inputFile + ".out"
	}

	err = ioutil.WriteFile(outputFile, []byte(encodedData), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		os.Exit(1)
	}

	fmt.Printf("File %s encoded and saved to %s\n", inputFile, outputFile)
}

func decodeFile(inputFile, outputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	decodedData, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		fmt.Println("Error decoding input file:", err)
		os.Exit(1)
	}

	if outputFile == "" {
		outputFile = inputFile + ".out"
	}

	err = ioutil.WriteFile(outputFile, []byte(decodedData), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		os.Exit(1)
	}

	fmt.Printf("File %s decoded and saved to %s\n", inputFile, outputFile)
}

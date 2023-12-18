package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type FileIO interface {
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

type realFileIO struct{}

func (r *realFileIO) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func (r *realFileIO) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}

var inputFile = flag.String("i", "", "Input file")
var outputFile = flag.String("o", "", "Output file")

var fileIO FileIO

func main() {
	fileIO = &realFileIO{}

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Usage: encoder64 [encode|decode] -i <inputfile> -o <outputfile>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "encode":
		err := flag.CommandLine.Parse(os.Args[2:])
		if err != nil {
			return
		}
		handleCommand("encode", *inputFile, *outputFile)

	case "decode":
		err := flag.CommandLine.Parse(os.Args[2:])
		if err != nil {
			return
		}
		handleCommand("decode", *inputFile, *outputFile)

	default:
		fmt.Println("Usage: encoder64 [encode|decode] -i <inputfile> -o <outputfile>")
		os.Exit(1)
	}
}

func handleCommand(command, inputFile, outputFile string) {
	switch command {
	case "encode":
		if inputFile == "" {
			fmt.Println("Input file is required for encoding.")
			os.Exit(1)
		}
		encodeFile(inputFile, outputFile)

	case "decode":
		if inputFile == "" {
			fmt.Println("Input file is required for decoding.")
			os.Exit(1)
		}
		decodeFile(inputFile, outputFile)
	}
}

func encodeFile(inputFile, outputFile string) {
	data, err := fileIO.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	encodedData := base64.StdEncoding.EncodeToString(data)

	if outputFile == "" {
		outputFile = inputFile + ".out"
	}

	err = fileIO.WriteFile(outputFile, []byte(encodedData), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		os.Exit(1)
	}

	fmt.Printf("File %s encoded and saved to %s\n", inputFile, outputFile)
}

func decodeFile(inputFile, outputFile string) {
	data, err := fileIO.ReadFile(inputFile)
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

	err = fileIO.WriteFile(outputFile, []byte(decodedData), 0644)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		os.Exit(1)
	}

	fmt.Printf("File %s decoded and saved to %s\n", inputFile, outputFile)
}

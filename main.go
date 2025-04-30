package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define command-line flags
	inputFiles := flag.String("i", "", "Comma-separated list of input CSV file paths")
	output := flag.String("o", "", "Path to the output CSV file (if omitted, writes to standard output)")
	flag.Parse()

	// Handle standard input if -i is not provided
	var mergedData [][]string
	if *inputFiles == "" {
		fmt.Println("Reading from standard input. Provide CSV data:")
		reader := csv.NewReader(os.Stdin)
		data, err := reader.ReadAll()
		if err != nil {
			fmt.Printf("Error reading from standard input: %v\n", err)
			os.Exit(1)
		}
		mergedData = data
	} else {
		// Parse input file paths
		filePaths := strings.Split(*inputFiles, ",")
		if len(filePaths) < 2 {
			fmt.Println("At least two input files must be specified.")
			os.Exit(1)
		}

		// Open and read all input files
		for idx, filePath := range filePaths {
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Printf("Error opening file %s: %v\n", filePath, err)
				os.Exit(1)
			}
			defer file.Close()

			reader := csv.NewReader(file)
			data, err := reader.ReadAll()
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", filePath, err)
				os.Exit(1)
			}

			if idx == 0 {
				mergedData = data
			} else {
				mergedData = mergeCSVData(mergedData, data)
			}
		}
	}

	// Write the merged data to the output file or standard output
	var writer *csv.Writer
	if *output == "" {
		writer = csv.NewWriter(os.Stdout)
	} else {
		outFile, err := os.Create(*output)
		if err != nil {
			fmt.Printf("Error creating output file: %v\n", err)
			os.Exit(1)
		}
		defer outFile.Close()
		writer = csv.NewWriter(outFile)
	}

	err := writer.WriteAll(mergedData)
	if err != nil {
		fmt.Printf("Error writing output: %v\n", err)
		os.Exit(1)
	}

	if *output != "" {
		fmt.Printf("Successfully merged files into %s\n", *output)
	}
}
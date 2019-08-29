package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	//Define variables
	var csvFile string
	var separator string
	var fileName string
	//Assuming source is in first column
	source := 0
	//Assuming destination is in second column
	destination := 1
	//Looking for user input for file path
	fmt.Println("Enter file path: (example /home/$USER/somefile.csv)")
	fmt.Scanln(&csvFile)
	fmt.Println("---------------------------------------------------")
	//Looking for user input for separator used in CSV file
	fmt.Println("Which separator you want to use?")
	fmt.Scanln(&separator)
	fmt.Println("---------------------------------------------------")
	fmt.Println("Name of output program ?")
	fmt.Scanln(&fileName)
	fmt.Println("---------------------------------------------------")
	//Creating file for config with input from fileName variable
	f, err := os.Create(fileName)
	//Opening CSV File.
	csvData, err := os.Open(csvFile)
	if err != nil {
		fmt.Println("Error: Can not open such a file")
	}
	r := csv.NewReader(csvData)
	//Generating Apache redirects
	fmt.Println("Apache redirects:")
	fmt.Println("-----------------")
	for {
		//Read CSV line
		reader, err := r.Read()
		//If end of file break loop
		if err == io.EOF {
			fmt.Println("Exited")
			break
		}
		//Handling error in reading file
		if err != nil {
			fmt.Println("Error reading file")
		}
		//Storing config syntax into variable
		comb := "\nlocation " + reader[source] + " { \n" + "rewrite ^/.* " + reader[destination] + "; \n" + "}"
		//Appending config syntax to file
		f, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
		_, err = f.WriteString(comb)
		f.Close()
	}
	fmt.Println("Complated - You can copy it to httpd.conf")
}

//WriteToFile Func is called in for loop to append values to file so it can be easy copied to httpd.conf
func WriteToFile(fileName string, data string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

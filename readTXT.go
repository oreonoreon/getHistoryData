package getHistoryData

import (
	"bufio"
	"log"
	"os"
)

// ReadTxT читает txt файл построчно, записывает результат в string slice. Принимает в качестве аргумента filename имя файла к примеру figi.txt
func readTxT(filename string) []string {
	var figis []string
	// open the file
	file, err := os.Open(filename)

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	// read line by line
	for fileScanner.Scan() {
		//fmt.Println(fileScanner.Text())
		figis = append(figis, fileScanner.Text())
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
	return figis
}

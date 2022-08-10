package getHistoryData

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func historyData(Figi string) {
	token := "t.yIDgyFK0WY0lmEQ76djHQRjNONyVBsB34gqdTmZrfdKTdhH0sz1nH9momdp1mCWn-mk7xLaQRym-xJXHDLTZYw"
	//Url := "https://invest-public-api.tinkoff.ru/history-data?figi=BBG000B9XRY4&year=2022"
	u := url.Values{}
	u.Add("figi", Figi)
	u.Add("year", "2022")
	Url := "https://invest-public-api.tinkoff.ru/history-data?" + u.Encode()
	client := &http.Client{}
	req, err := http.NewRequest("GET", Url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	file, err := os.Create(Figi + ".zip")
	if err != nil {
		panic(err)
	}
	// Получить объект записи файла
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("Total length: %d", written)

}

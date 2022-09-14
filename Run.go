package getHistoryData

// Run запускает процесс запроса тинкофф и разиповки в файлы filename - имя txt файла где записаны figi инстументов, Token - токен писочницы Тинькоф
func Run(filename string, Token string) {
	figis := readTxT(filename)
	for _, value := range figis {
		historyData(value, Token)
		unzip(value + ".zip")
	}
}

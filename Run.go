package getHistoryData

// Run запускает процесс запроса тинкофф и разиповки в файлы
func Run(filename string) {
	figis := readTxT(filename)
	for _, value := range figis {
		historyData(value)
		unzip(value + ".zip")
	}
}

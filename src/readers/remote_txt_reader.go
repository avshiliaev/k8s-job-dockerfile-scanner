package readers

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"redhat-sre-task-dockerfile-scanner/src/models"
)

type remoteTxtReader struct {
	client HttpClient
}

func RemoteTxtReader(client HttpClient) *remoteTxtReader {
	return &remoteTxtReader{
		client: client,
	}
}

func (reader *remoteTxtReader) Read(data *models.Data) {

	var inputLines []string
	req, err := http.NewRequest("GET", data.Url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Content-Type: text/plain
	// Content-Disposition: attachment; filename="yourfilename.txt"
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Content-Disposition", "attachment; filename=\"sources.txt\"")

	resp, err := reader.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	fileReader := bufio.NewScanner(resp.Body)

	for fileReader.Scan() {
		input := fileReader.Text()
		// TODO: checks
		inputLines = append(inputLines, input)
	}
	if err := fileReader.Err(); err != nil {
		log.Fatal(err)
	}

	data.InputLines = inputLines

}

package readers

import (
	"bufio"
	"io"
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

func (reader *remoteTxtReader) Read(data *models.Data) error {

	var err error
	var inputLines []string
	req, err := http.NewRequest("GET", data.Url, nil)
	if err != nil {
		return err
	}

	// Content-Type: text/plain
	// Content-Disposition: attachment; filename="yourfilename.txt"
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Content-Disposition", "attachment; filename=\"sources.txt\"")

	resp, err := reader.client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) error {
		err := Body.Close()
		if err != nil {
			return err
		}
		return nil
	}(resp.Body)

	fileReader := bufio.NewScanner(resp.Body)

	for fileReader.Scan() {
		input := fileReader.Text()
		inputLines = append(inputLines, input)
	}
	if err := fileReader.Err(); err != nil {
		return err
	}

	data.InputLines = inputLines
	return err

}

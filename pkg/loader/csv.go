package loader

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

type Loader interface {
	ReadCSV(reader io.ReadCloser, separator rune, skipHeader bool) (chan Record, error)
	DownloadFile(url string) (io.ReadCloser, error)
}

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type Record struct {
	Product Product
	Err     error
}

type CSVLoader struct {
	logger *logrus.Entry
}

func (l CSVLoader) DownloadFile(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url) // nolint
	if err != nil {
		return nil, fmt.Errorf("cannot read url: %s, %w", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad http response status: %s", resp.Status)
	}
	return resp.Body, nil
}

func (l CSVLoader) ReadCSV(reader io.ReadCloser, separator rune, skipHeader bool) (chan Record, error) {
	csvReader := csv.NewReader(reader)
	csvReader.Comma = separator
	if skipHeader {
		_, err := csvReader.Read()
		if err != nil {
			return nil, err
		}
	}
	records := make(chan Record)
	go func() {
		defer func() {
			close(records)
			if err := reader.Close(); err != nil {
				l.logger.Errorf("cannot close http body: %v", err)
			}
		}()
		for {
			csvRecord, err := csvReader.Read()
			if errors.Is(err, io.EOF) {
				break
			}
			if err != nil {
				records <- Record{Err: err}
				break
			}
			if len(csvRecord) < 2 {
				records <- Record{Err: fmt.Errorf("malformed CSV record: %s", csvRecord)}
				break
			}
			price, err := strconv.Atoi(csvRecord[1])
			if err != nil {
				records <- Record{Err: fmt.Errorf("malformed CSV record %s: %w", csvRecord, err)}
				break
			}
			record := Product{
				Name:  csvRecord[0],
				Price: price,
			}

			records <- Record{Product: record}
		}
	}()

	return records, nil
}

func NewCsvLoader(logger *logrus.Entry) *CSVLoader {
	return &CSVLoader{logger: logger}
}

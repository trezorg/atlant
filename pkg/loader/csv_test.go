package loader

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sirupsen/logrus"
)

const csvData = `test1;100
test2;102
test3;103`

func TestLoader(t *testing.T) {
	loader := NewCsvLoader(logrus.WithFields(logrus.Fields{"service": "atlant-grpc"}))
	out, err := loader.ReadCSV(ioutil.NopCloser(strings.NewReader(csvData)), ';', false)
	require.NoError(t, err)
	var result []Record // nolint
	for rec := range out {
		result = append(result, rec)
	}
	require.Len(t, result, 3)
}

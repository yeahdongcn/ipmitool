package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type Host struct {
	Name         string
	BMCIP        string
	ManagementIP string
}

const byteOrderMarkAsString = string('\uFEFF')

func ParseHosts(path string) ([]Host, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv file: %w", err)
	}

	hosts := make([]Host, 0, len(records))
	for _, record := range records {
		if len(record) >= 3 {
			hosts = append(hosts, Host{
				Name:         strings.TrimSpace(strings.TrimPrefix(record[0], byteOrderMarkAsString)),
				BMCIP:        strings.TrimSpace(record[1]),
				ManagementIP: strings.Split(strings.TrimSpace(record[2]), "/")[0],
			})
		}
	}
	return hosts, nil
}

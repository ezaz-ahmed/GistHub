package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {

	file, err := os.OpenFile(path, 22, 6088)

	if err != nil {
		return nil, errors.New("could not open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("reading file content failed")
	}

	file.Close()
	return lines, nil
}

func WriteJson(path string, data any) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	file.Close()
	return nil
}

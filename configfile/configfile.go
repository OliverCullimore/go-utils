package configfile

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// Save function
func Save(filename string, conf interface{}) error {
	bytes, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, bytes, 0644)
}

// Load function
func Load(filename string) (interface{}, error) {
	// Check file exists
	info, err := os.Stat(filename)
	if os.IsNotExist(err) || info.IsDir() {
		return nil, errors.New("file not found")
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error loading file")
	}

	var configuration interface{}
	err = json.Unmarshal(bytes, &configuration)
	if err != nil {
		return nil, err
	}

	return configuration, nil
}

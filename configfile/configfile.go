package configfile

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// Save function
func Save(filename string, config interface{}) error {
	// Parse config interface into JSON
	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	// Write JSON to file
	return ioutil.WriteFile(filename, bytes, 0644)
}

// Load function
func Load(filename string, config interface{}) error {
	// Check file exists
	info, err := os.Stat(filename)
	if os.IsNotExist(err) || info.IsDir() {
		return errors.New("file not found")
	}
	// Read JSON file contents
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	// Parse JSON into config interface
	err = json.Unmarshal(bytes, config)
	if err != nil {
		return err
	}
	return nil
}

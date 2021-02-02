package configfile

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// Save function accepts a filename and a config interface
// to parse the config data from
func Save(filename string, config interface{}) error {
	// Parse config interface into JSON
	bytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	// Write JSON to file
	return ioutil.WriteFile(filename, bytes, 0644)
}

// Load function accepts a filename and a config interface
// to parse the config data into
func Load(filename string, config interface{}) error {
	// Check file exists
	info, err := os.Stat(filename)
	if os.IsNotExist(err) || info.IsDir() {
		return errors.New(fmt.Sprintf("config file %s not found", filename))
	}
	if err == nil {
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
	}
	return nil
}

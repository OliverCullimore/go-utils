package env

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

var envVars map[string]string

// Load checks if the specified environment variables file exists and loads it if it does,
// if it doesn't an error will be returned instead.
func Load(envFile string) error {
	// Check file exists
	info, err := os.Stat(envFile)
	if os.IsNotExist(err) || info.IsDir() {
		return errors.New("file not found")
	}
	envVars, err = godotenv.Read()
	if err != nil {
		return errors.New("error loading file")
	}
	return nil
}

// Get checks if the specified environment variable key exists and returns the value,
// if it doesn't the specified default value will be returned instead.
func Get(envKey, defaultVal string) string {
	// Check for an environment variable
	envVal, envPresent := os.LookupEnv(envKey)
	if envPresent && envVal != "" {
		return envVal
	}
	// Check the loaded vars
	if val, ok := envVars[envKey]; ok && val != "" {
		return val
	}
	return defaultVal
}

/*
// Set sets the value of the environment variable.
func Set(envKey, envVal string) error {
	return os.Setenv(envKey, envVal)
}

// Unset unsets a single environment variable.
func Unset(envKey string) error {
	return os.Unsetenv(envKey)
}

// Clear deletes all environment variables.
func Clear() {
	os.Clearenv()
}
*/

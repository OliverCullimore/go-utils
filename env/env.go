package env

import (
	"os"
)

// Get checks if the specified environment variable key exists and returns the value,
// if it doesn't the specified default value will be returned instead.
func Get(envKey, defaultVal string) string {
	envVal, envPresent := os.LookupEnv(envKey)
	if envPresent && envVal != "" {
		return envVal
	}
	return defaultVal
}

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

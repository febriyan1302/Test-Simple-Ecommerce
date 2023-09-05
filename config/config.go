package config

import (
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
)

// reStringFromEnv regex format for string from environment variable.
var reStringFromEnv = regexp.MustCompile(`^\${(\w+)}$`)

type stringFromEnv string

// UnmarshalYAML decode environtment variable value from configuration YAML file.
// Accept pointer of yaml.Node.
// It return nil error if succeed, otherwise return non-nil error.
func (e *stringFromEnv) UnmarshalYAML(value *yaml.Node) error {
	var s string
	if err := value.Decode(&s); err != nil {
		return err
	}

	if match := reStringFromEnv.FindStringSubmatch(s); len(match) > 0 {
		*e = stringFromEnv(os.Getenv(match[1]))
	} else {
		*e = stringFromEnv(s)
	}

	return nil
}

type Config struct {
	App struct {
		Name     stringFromEnv `yaml:"name"`
		Env      string        `yaml:"env"`
		RestPort stringFromEnv `yaml:"rest_port"`
	} `yaml:"app"`

	Database struct {
		Host     stringFromEnv `yaml:"host"`
		Port     stringFromEnv `yaml:"port"`
		Name     stringFromEnv `yaml:"name"`
		User     stringFromEnv `yaml:"user"`
		Password stringFromEnv `yaml:"password"`
	} `yaml:"db"`
}

// Обработка конфигурации пакета
// без проверок зависимостей, только стат анализ.

package info

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"unicode/utf8"

	"github.com/ghodss/yaml"
)

type Config struct {
	Version      string            `yaml:"version" json:"version"`
	Name         string            `yaml:"name" json:"name"`
	Region       string            `yaml:"region" json:"region"`
	Description  string            `yaml:"description,omitempty" json:"description,omitempty"`
	Tags         []string          `yaml:"tags,omitempty" json:"tags,omitempty"`
	Dependencies map[string]string `yaml:"dependencies,omitempty" json:"dependencies,omitempty"`
}

// чтение + валидация конфигурации пакета
func ReadConfig(dir string) (*Config, error) {
	configPath := filepath.Join(dir, "package.yml")

	config, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read package.yml: %w", err)
	}

	var pkg Config
	if err := yaml.Unmarshal(config, &pkg); err != nil {
		return nil, fmt.Errorf("invalid YAML in package.yml: %w", err)
	}

	if err := validateConfig(pkg); err != nil {
		return nil, fmt.Errorf("error validating package.yml: %w", err)
	}

	return &pkg, nil
}

// валидация конвига
func validateConfig(conf Config) error {
	// name
	if conf.Name == "" {
		return fmt.Errorf("'name' is required")
	}

	nameLen := utf8.RuneCountInString(conf.Name)
	if nameLen < 4 || nameLen > 32 {
		return fmt.Errorf(
			"'name' must be 4-32 characters, got %d (%s)",
			nameLen, conf.Name)
	}

	// version
	if conf.Version == "" {
		return fmt.Errorf("'version' is required")
	}

	versionRegex := regexp.MustCompile(`^\d+\.\d+\.\d+$`)
	if !versionRegex.MatchString(conf.Version) {
		return fmt.Errorf(
			"'version' must be X.Y.Z format (e.g. 0.0.1), got %s",
			conf.Version)
	}

	// region
	if conf.Region != "" {
		regionLen := utf8.RuneCountInString(conf.Region)
		if regionLen < 2 || regionLen > 3 {
			return fmt.Errorf(
				"'region' must be 2-3 characters (e.g. 'ru'), got %s",
				conf.Region)
		}
	}

	return nil
}

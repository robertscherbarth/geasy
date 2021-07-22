package printer

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/pkg/errors"
	"github.com/robertscherbarth/geasy/internal/flags"
)

func Warnf(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}

func Errorf(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}

func Infof(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}

func Debugf(msg string, args ...interface{}) {
	if flags.Verbose {
		fmt.Printf(msg+"\n", args...)
	}
}

func PrintObject(object interface{}, format string) error {
	switch format {
	case "yaml", "yml":
		yamlString, err := yaml.Marshal(object)
		if err != nil {
			return fmt.Errorf("could not format yaml cause of: %w", err)
		}
		fmt.Println(string(yamlString))
		return nil
	case "json":
		jsonString, err := json.Marshal(object)
		if err != nil {
			return fmt.Errorf("could not format yaml cause of: %w", err)
		}
		fmt.Println(string(jsonString))
		return nil
	default:
		return errors.Errorf("unknown format: %v", format)
	}
}

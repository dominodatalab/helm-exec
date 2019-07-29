package helmexec

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type Plugin struct {
	Name        string
	Version     string
	Description string
}

func (w Wrapper) PluginInstall(pathOrURL, version string) error {
	if pathOrURL == "" {
		return errors.New("pathOrURL cannot be empty")
	}

	args := []string{"plugin", "install", pathOrURL}
	if version != "" {
		args = append(args, "--version", version)
	}

	_, err := w.exec(args...)
	return err
}

func (w Wrapper) PluginList() (plugins []Plugin, err error) {
	out, err := w.exec("plugin", "list")
	if err != nil {
		return
	}

	re := regexp.MustCompile(`(?m)^(\S+)\s+(\S+)\s+(.*)$`)
	matches := re.FindAllStringSubmatch(string(out), -1)
	header := []string{"NAME", "VERSION", "DESCRIPTION"}
	if matches == nil || !reflect.DeepEqual(strings.Fields(matches[0][0]), header) {
		return nil, fmt.Errorf("invalid list format: %s", string(out))
	}

	for _, data := range matches[1:] {
		plugin := Plugin{
			Name:        data[1],
			Version:     data[2],
			Description: data[3],
		}
		plugins = append(plugins, plugin)
	}
	return
}

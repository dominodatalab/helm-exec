package helmexec

import (
	"errors"
	"regexp"
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
	if matches == nil {
		return
	}

	for _, data := range matches {
		if data[1] == "NAME" { // ignore header
			continue
		}
		plugin := Plugin{
			Name:        data[1],
			Version:     data[2],
			Description: data[3],
		}
		plugins = append(plugins, plugin)
	}
	return
}

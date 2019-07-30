package helmexec

import (
	"errors"
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
	cmdArgs := []string{"plugin", "list"}
	headerOutput := []string{"NAME", "VERSION", "DESCRIPTION"}
	captureRegex := `(?m)^(\S+)\s+(\S+)\s+(.*)$`
	err = w.listAction(cmdArgs, headerOutput, captureRegex, func(item []string) {
		plugin := Plugin{
			Name:        item[1],
			Version:     item[2],
			Description: item[3],
		}
		plugins = append(plugins, plugin)
	})
	return
}

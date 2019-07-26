package helmexec

import "errors"

type Plugin struct {
	Name        string
	Version     string
	Description string
}

func (w wrapper) PluginInstall(pathOrURL, version string) error {
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

func (w wrapper) PluginList() ([]Plugin, error) {
	return nil, nil
}

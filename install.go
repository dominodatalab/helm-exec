package helmexec

import "fmt"

type InstallOptions struct {
	Name        string
	Namespace   string
	Description string
	Version     string
	Wait        bool
	Set         map[string]string
}

func (w Wrapper) Install(chstr string, opts *InstallOptions) error {
	args := []string{"install", chstr}
	if opts != nil {
		if opts.Name != "" {
			args = append(args, fmt.Sprintf("--name=%s", opts.Name))
		}
		if opts.Namespace != "" {
			args = append(args, fmt.Sprintf("--namespace=%s", opts.Namespace))
		}
		if opts.Description != "" {
			args = append(args, fmt.Sprintf("--description=%q", opts.Description))
		}
		if opts.Version != "" {
			args = append(args, fmt.Sprintf("--version=%s", opts.Version))
		}
		if opts.Wait {
			args = append(args, fmt.Sprintf("--wait"))
		}
		for key, value := range opts.Set {
			args = append(args, "--set", fmt.Sprintf("%s=%q", key, value))
		}
	}

	_, err := w.exec(args...)
	return err
}

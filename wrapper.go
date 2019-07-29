package helmexec

import (
	"errors"

	"github.com/dominodatalab/vagrant-exec/command"
)

const binary = "helm"

type Wrapper struct {
	executable string
	runner     command.Runner
}

func New() Wrapper {
	return Wrapper{
		executable: binary,
		runner:     command.ShellRunner{},
	}
}

func (w *Wrapper) SetRunner(runner command.Runner) error {
	if runner != nil {
		w.runner = runner
		return nil
	}
	return errors.New("command runner cannot be nil")
}

func (w Wrapper) exec(args ...string) ([]byte, error) {
	return w.runner.Execute(w.executable, args...)
}

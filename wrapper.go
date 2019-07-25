package helmexec

import "github.com/dominodatalab/vagrant-exec/command"

const binary = "helm"

type wrapper struct {
	executable string
	runner     command.Runner
}

func New() wrapper {
	return wrapper{
		executable: binary,
		runner:     command.ShellRunner{},
	}
}

func (w wrapper) exec(args ...string) ([]byte, error) {
	return w.runner.Execute(w.executable, args...)
}

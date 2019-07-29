package helmexec_test

import (
	"testing"

	he "github.com/dominodatalab/helm-exec"
	"github.com/dominodatalab/vagrant-exec/command"
	"github.com/stretchr/testify/assert"
)

type FakeRunner struct {
	execFn func(cmd []string) ([]byte, error)
}

func (m *FakeRunner) Execute(cmd string, args ...string) ([]byte, error) {
	fullCmd := append([]string{cmd}, args...)
	return m.execFn(fullCmd)
}

func NewTestWrapper() (he.Wrapper, *FakeRunner) {
	runner := new(FakeRunner)
	runner.execFn = func([]string) ([]byte, error) {
		return nil, nil
	}

	helm := he.New()
	err := helm.SetRunner(runner)
	if err != nil {
		panic(err)
	}

	return helm, runner
}

func TestWrapper_SetRunner(t *testing.T) {
	helm := he.Wrapper{}

	assert.Error(t, helm.SetRunner(nil))
	assert.NoError(t, helm.SetRunner(&command.ShellRunner{}))
}

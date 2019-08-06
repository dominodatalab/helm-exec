package helmexec_test

import (
	"errors"
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

func assertRunnerErr(t *testing.T, runner *FakeRunner, fn func() error) {
	t.Helper()

	errMsg := "runner error"
	runner.execFn = func([]string) ([]byte, error) {
		return nil, errors.New(errMsg)
	}
	assert.EqualError(t, fn(), errMsg)
}

func TestWrapper_Delete(t *testing.T) {
	helm, runner := NewTestWrapper()

	delFnSuccess := func(name string) func([]string) ([]byte, error) {
		return func(cmd []string) ([]byte, error) {
			assert.Equal(t, []string{"helm", "delete", name}, cmd)
			return nil, nil
		}
	}

	t.Run("Successful Delete", func(t *testing.T) {
		release := "phpmyadmin"
		runner.execFn = delFnSuccess(release)
		assert.NoError(t, helm.Delete(release))
	})

	delFnFailure := func(name string) func([]string) ([]byte, error) {
		return func(cmd []string) ([]byte, error) {
			assert.Equal(t, []string{"helm", "delete", name}, cmd)
			return nil, errors.New("Helm delete failure")
		}
	}

	t.Run("Errored Delete", func(t *testing.T) {
		release := "sin"
		runner.execFn = delFnFailure(release)
		assert.Error(t, helm.Delete(release))
	})

}

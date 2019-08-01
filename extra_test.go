package helmexec_test

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapper_IsRelease(t *testing.T) {
	helm, runner := NewTestWrapper()

	rlsFn := func(name string) func([]string) ([]byte, error) {
		return func(cmd []string) ([]byte, error) {
			assert.Equal(t, []string{"helm", "list", "--all", "--output=json", name}, cmd)
			return ioutil.ReadFile("testdata/release-list-one")
		}
	}

	t.Run("exists", func(t *testing.T) {
		rlsName := "postgresql"
		runner.execFn = rlsFn(rlsName)
		assert.True(t, helm.IsRelease(rlsName))
	})

	t.Run("missing", func(t *testing.T) {
		missing := "missing-rls"
		runner.execFn = rlsFn(missing)
		assert.False(t, helm.IsRelease(missing))
	})

	t.Run("error", func(t *testing.T) {
		runner.execFn = func([]string) ([]byte, error) {
			return nil, errors.New("runner error")
		}
		assert.False(t, helm.IsRelease("notarealthing"))
	})
}

func TestWrapper_IsRepo(t *testing.T) {
}

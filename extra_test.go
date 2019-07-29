package helmexec_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapper_IsRelease(t *testing.T) {
	helm, runner := NewTestWrapper()

	rlsName := "some-rls"
	rlsFn := func(name string) func([]string) ([]byte, error) {
		return func(cmd []string) ([]byte, error) {
			assert.Equal(t, []string{"helm", "list", "--all", "--short", name}, cmd)
			return []byte(fmt.Sprintf("%s\nanother-rls\n", rlsName)), nil
		}
	}

	t.Run("exists", func(t *testing.T) {
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
		assert.False(t, helm.IsRelease(rlsName))
	})
}

func TestWrapper_IsRepo(t *testing.T) {
}

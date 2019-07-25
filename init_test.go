package helmexec

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapper_Init(t *testing.T) {
	runner := new(FakeRunner)
	helm := New()
	helm.runner = runner

	t.Run("success", func(t *testing.T) {
		runner.execFn = func(cmd []string) (bytes []byte, e error) {
			assert.Equal(t, []string{"helm", "init"}, cmd)
			return nil, nil
		}

		assert.NoError(t, helm.Init())
	})

	t.Run("upgrade", func(t *testing.T) {
		runner.execFn = func(cmd []string) (bytes []byte, e error) {
			assert.Equal(t, []string{"helm", "init", "--upgrade"}, cmd)
			return nil, nil
		}

		assert.NoError(t, helm.Init(InitUpgrade(true)))
	})

	t.Run("wait", func(t *testing.T) {
		runner.execFn = func(cmd []string) (bytes []byte, e error) {
			assert.Equal(t, []string{"helm", "init", "--wait"}, cmd)
			return nil, nil
		}

		assert.NoError(t, helm.Init(InitWait(true)))
	})

	t.Run("error", func(t *testing.T) {
		runner.execFn = func(cmd []string) (bytes []byte, e error) {
			return nil, errors.New("runner error")
		}

		assert.Error(t, helm.Init())
	})
}

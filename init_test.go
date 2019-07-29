package helmexec_test

import (
	"errors"
	"strings"
	"testing"

	he "github.com/dominodatalab/helm-exec"
	"github.com/stretchr/testify/assert"
)

func TestWrapper_Init(t *testing.T) {
	helm, runner := NewTestWrapper()

	t.Run("success", func(t *testing.T) {
		runner.execFn = func(cmd []string) ([]byte, error) {
			assert.Equal(t, strings.Fields("helm init"), cmd)
			return nil, nil
		}
		assert.NoError(t, helm.Init())
	})

	t.Run("upgrade", func(t *testing.T) {
		runner.execFn = func(cmd []string) ([]byte, error) {
			assert.Equal(t, strings.Fields("helm init --upgrade"), cmd)
			return nil, nil
		}
		assert.NoError(t, helm.Init(he.InitUpgrade(true)))
	})

	t.Run("wait", func(t *testing.T) {
		runner.execFn = func(cmd []string) ([]byte, error) {
			assert.Equal(t, strings.Fields("helm init --wait"), cmd)
			return nil, nil
		}
		assert.NoError(t, helm.Init(he.InitWait(true)))
	})

	t.Run("error", func(t *testing.T) {
		runner.execFn = func(cmd []string) ([]byte, error) {
			return nil, errors.New("runner error")
		}
		assert.EqualError(t, helm.Init(), "runner error")
	})
}

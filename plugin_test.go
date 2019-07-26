package helmexec

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapper_PluginInstall(t *testing.T) {
	runner := new(FakeRunner)
	helm := New()
	helm.runner = runner

	t.Run("success", func(t *testing.T) {
		runner.execFn = func(cmd []string) (bytes []byte, e error) {
			exp := strings.Fields("helm plugin install path-or-url")
			assert.Equal(t, exp, cmd)

			return nil, nil
		}
		assert.NoError(t, helm.PluginInstall("path-or-url", ""))
	})

	t.Run("with_version", func(t *testing.T) {
		runner.execFn = func(cmd []string) ([]byte, error) {
			exp := strings.Fields("helm plugin install path-or-url --version 1.2.3")
			assert.Equal(t, exp, cmd)

			return nil, nil
		}
		assert.NoError(t, helm.PluginInstall("path-or-url", "1.2.3"))
	})

	t.Run("empty_url", func(t *testing.T) {
		runner.execFn = func(cmd []string) (bytes []byte, e error) { return nil, nil }
		err := helm.PluginInstall("", "")

		assert.EqualError(t, err, "pathOrURL cannot be empty")
	})

	t.Run("error", func(t *testing.T) {
		runner.execFn = func(cmd []string) (bytes []byte, e error) {
			return nil, errors.New("runner error")
		}

		assert.Error(t, helm.PluginInstall("plugin-location", "1.2.3"))
	})
}

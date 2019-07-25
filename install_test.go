package helmexec

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapper_Install(t *testing.T) {
	runner := new(FakeRunner)
	helm := New()
	helm.runner = runner

	t.Run("no_opts", func(t *testing.T) {
		runner.execFn = func(cmd []string) (bytes []byte, e error) {
			assert.Equal(t, []string{"helm", "install", "chart-url-or-path"}, cmd)
			return nil, nil
		}
		assert.NoError(t, helm.Install("chart-url-or-path", nil))
	})

	t.Run("full_opts", func(t *testing.T) {
		opts := &InstallOptions{
			Name:        "rls",
			Namespace:   "my-ns",
			Description: "awesome opossum",
			Version:     "1.2.3",
			Wait:        true,
			Set: map[string]string{
				"one":    "two",
				"buckle": "my shoe",
			},
		}
		runner.execFn = func(cmd []string) (bytes []byte, e error) {
			fullCmd := []string{
				"helm", "install", "chart-url-or-path",
				"--name=rls",
				"--namespace=my-ns",
				`--description="awesome opossum"`,
				"--version=1.2.3",
				"--wait",
				"--set", `one="two"`,
				"--set", `buckle="my shoe"`,
			}
			assert.Equal(t, fullCmd, cmd)
			return nil, nil
		}
		assert.NoError(t, helm.Install("chart-url-or-path", opts))
	})

	t.Run("error", func(t *testing.T) {
		runner.execFn = func(cmd []string) (bytes []byte, e error) {
			return nil, errors.New("runner error")
		}
		assert.Error(t, helm.Install("chart-url-or-path", nil))
	})
}

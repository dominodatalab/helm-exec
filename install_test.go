package helmexec_test

import (
	"testing"

	he "github.com/dominodatalab/helm-exec"
	"github.com/stretchr/testify/assert"
)

func TestWrapper_Install(t *testing.T) {
	helm, runner := NewTestWrapper()
	chartStr := "chart-url-or-path"

	t.Run("no_opts", func(t *testing.T) {
		runner.execFn = func(cmd []string) ([]byte, error) {
			assert.Equal(t, []string{"helm", "install", chartStr}, cmd)
			return nil, nil
		}
		assert.NoError(t, helm.Install(chartStr, nil))
	})

	t.Run("full_opts", func(t *testing.T) {
		opts := &he.InstallOptions{
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
		runner.execFn = func(cmd []string) ([]byte, error) {
			leadingCmd := cmd[:len(cmd)-4]
			exp := []string{
				"helm", "install", chartStr,
				"--name=rls",
				"--namespace=my-ns",
				`--description="awesome opossum"`,
				"--version=1.2.3",
				"--wait",
			}
			assert.Equal(t, exp, leadingCmd)

			setFlags := cmd[len(cmd)-4:]
			exp = []string{
				"--set", `one="two"`,
				"--set", `buckle="my shoe"`,
			}
			assert.ElementsMatch(t, exp, setFlags)

			return nil, nil
		}
		assert.NoError(t, helm.Install(chartStr, opts))
	})

	assertRunnerErr(t, runner, func() error {
		return helm.Install(chartStr, nil)
	})
}

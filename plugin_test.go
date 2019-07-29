package helmexec_test

import (
	"errors"
	"io/ioutil"
	"testing"

	he "github.com/dominodatalab/helm-exec"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrapper_PluginInstall(t *testing.T) {
	helm, runner := NewTestWrapper()
	pluginName := "path-or-url"
	pluginVersion := "1.2.3"

	t.Run("success", func(t *testing.T) {
		runner.execFn = func(cmd []string) ([]byte, error) {
			assert.Equal(t, []string{"helm", "plugin", "install", pluginName}, cmd)
			return nil, nil
		}
		assert.NoError(t, helm.PluginInstall(pluginName, ""))
	})

	t.Run("with_version", func(t *testing.T) {
		runner.execFn = func(cmd []string) ([]byte, error) {
			assert.Equal(t, []string{"helm", "plugin", "install", pluginName, "--version", pluginVersion}, cmd)
			return nil, nil
		}
		assert.NoError(t, helm.PluginInstall(pluginName, pluginVersion))
	})

	t.Run("empty_url", func(t *testing.T) {
		runner.execFn = func([]string) ([]byte, error) { return nil, nil }
		err := helm.PluginInstall("", "")
		assert.EqualError(t, err, "pathOrURL cannot be empty")
	})

	t.Run("error", func(t *testing.T) {
		runner.execFn = func([]string) ([]byte, error) {
			return nil, errors.New("runner error")
		}
		assert.EqualError(t, helm.PluginInstall(pluginName, pluginVersion), "runner error")
	})
}

func TestWrapper_PluginList(t *testing.T) {
	helm, runner := NewTestWrapper()

	testcases := []struct {
		name    string
		fixture string
		exp     []he.Plugin
	}{
		{
			"no_plugins",
			"testdata/plugin-list-none",
			nil,
		},
		{
			"one_plugin",
			"testdata/plugin-list-one",
			[]he.Plugin{
				{Name: "registry", Version: "0.7.0", Description: "This plugin provides app-registry client to Helm."},
			},
		},
		{
			"many_plugins",
			"testdata/plugin-list-many",
			[]he.Plugin{
				{Name: "registry", Version: "0.7.0", Description: "This plugin provides app-registry client to Helm."},
				{Name: "diff", Version: "2.11.0+5", Description: "Preview helm upgrade changes as a diff"},
				{Name: "env", Version: "0.1.0", Description: "Print out the helm environment."},
			},
		},
		{
			"unexpected_format",
			"testdata/plugin-list-unexpected",
			nil,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			runner.execFn = func([]string) ([]byte, error) {
				return ioutil.ReadFile(tc.fixture)
			}
			list, err := helm.PluginList()
			require.NoError(t, err)
			assert.Equal(t, tc.exp, list)
		})
	}

	t.Run("error", func(t *testing.T) {
		runner.execFn = func([]string) ([]byte, error) {
			return nil, errors.New("runner error")
		}
		_, err := helm.PluginList()
		assert.EqualError(t, err, "runner error")
	})
}

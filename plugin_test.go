package helmexec

import (
	"errors"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestWrapper_PluginList(t *testing.T) {
	runner := new(FakeRunner)
	helm := New()
	helm.runner = runner

	testcases := []struct {
		name    string
		fixture string
		exp     []Plugin
	}{
		{
			"no_plugins",
			"testdata/plugin-list-none",
			nil,
		},
		{
			"one_plugin",
			"testdata/plugin-list-one",
			[]Plugin{
				{Name: "registry", Version: "0.7.0", Description: "This plugin provides app-registry client to Helm."},
			},
		},
		{
			"many_plugins",
			"testdata/plugin-list-many",
			[]Plugin{
				{Name: "registry", Version: "0.7.0", Description: "This plugin provides app-registry client to Helm."},
				{Name: "diff", Version: "2.11.0+5", Description: "Preview helm upgrade changes as a diff"},
				{Name: "env", Version: "0.1.0", Description: "Print out the helm environment."},
			},
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
		runner.execFn = func(cmd []string) (bytes []byte, e error) {
			return nil, errors.New("runner error")
		}

		_, err := helm.PluginList()
		assert.EqualError(t, err, "runner error")
	})
}

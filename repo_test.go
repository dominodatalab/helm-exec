package helmexec_test

import (
	"io/ioutil"
	"testing"

	he "github.com/dominodatalab/helm-exec"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWrapper_RepoAdd(t *testing.T) {
	helm, runner := NewTestWrapper()
	repoName := "domino"
	repoURL := "https://charts.dominodatalab.com"

	t.Run("required_args", func(t *testing.T) {
		testcases := []struct {
			name, url string
		}{
			{repoName, ""},
			{"", repoURL},
			{"", ""},
		}
		for _, tc := range testcases {
			err := helm.RepoAdd(tc.name, tc.url, nil)
			assert.EqualError(t, err, "repo name and url are required")
		}

		runner.execFn = func(cmd []string) ([]byte, error) {
			exp := []string{"helm", "repo", "add", repoName, repoURL}
			assert.Equal(t, exp, cmd)
			return nil, nil
		}
		assert.NoError(t, helm.RepoAdd(repoName, repoURL, nil))
	})

	t.Run("options", func(t *testing.T) {
		testcases := []struct {
			opts *he.RepoAddOptions
			exp  []string
		}{
			{
				&he.RepoAddOptions{},
				[]string{"helm", "repo", "add", repoName, repoURL},
			},
			{
				&he.RepoAddOptions{
					NoUpdate: true,
				},
				[]string{"helm", "repo", "add", repoName, repoURL, "--no-update"},
			},
			{
				&he.RepoAddOptions{
					Username: "ewe",
					Password: "know",
				},
				[]string{"helm", "repo", "add", repoName, repoURL, "--username", "ewe", "--password", "know"},
			},
			{
				&he.RepoAddOptions{
					Username: "missing-pass",
				},
				[]string{"helm", "repo", "add", repoName, repoURL},
			},
			{
				&he.RepoAddOptions{
					Password: "missing-user",
				},
				[]string{"helm", "repo", "add", repoName, repoURL},
			},
		}
		for _, tc := range testcases {
			runner.execFn = func(cmd []string) ([]byte, error) {
				assert.Equal(t, tc.exp, cmd)
				return nil, nil
			}
			assert.NoError(t, helm.RepoAdd(repoName, repoURL, tc.opts))
		}
	})

	assertRunnerErr(t, runner, func() error {
		return helm.RepoAdd(repoName, repoURL, nil)
	})
}

func TestWrapper_RepoList(t *testing.T) {
	helm, runner := NewTestWrapper()

	testcases := []struct {
		name    string
		fixture string
		exp     []he.Repository
	}{
		{
			"one_repo",
			"testdata/repo-list-one",
			[]he.Repository{
				{Name: "local", URL: "http://127.0.0.1:8879/charts"},
			},
		},
		{
			"many_repos",
			"testdata/repo-list-many",
			[]he.Repository{
				{Name: "stable", URL: "https://kubernetes-charts.storage.googleapis.com"},
				{Name: "local", URL: "http://127.0.0.1:8879/charts"},
				{Name: "rancher-stable", URL: "https://releases.rancher.com/server-charts/stable"},
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			runner.execFn = func([]string) ([]byte, error) {
				return ioutil.ReadFile(tc.fixture)
			}
			list, err := helm.RepoList()
			require.NoError(t, err)
			assert.Equal(t, tc.exp, list)
		})
	}

	t.Run("bad_format", func(t *testing.T) {
		runner.execFn = func([]string) ([]byte, error) {
			return []byte("some unexpected repo list format"), nil
		}
		_, err := helm.RepoList()
		assert.Error(t, err)
	})

	assertRunnerErr(t, runner, func() error {
		_, err := helm.RepoList()
		return err
	})
}

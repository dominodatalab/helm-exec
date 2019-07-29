package helmexec_test

import (
	"errors"
	"testing"

	he "github.com/dominodatalab/helm-exec"
	"github.com/stretchr/testify/assert"
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

	t.Run("error", func(t *testing.T) {
		errMsg := "runner error"
		runner.execFn = func([]string) ([]byte, error) {
			return nil, errors.New(errMsg)
		}
		assert.EqualError(t, helm.RepoAdd(repoName, repoURL, nil), errMsg)
	})
}

func TestWrapper_RepoList(t *testing.T) {
}

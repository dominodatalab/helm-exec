package helmexec

import "errors"

type RepoAddOptions struct {
	Username string
	Password string
	NoUpdate bool
}

type Repository struct {
	Name string
	URL  string
}

func (w Wrapper) RepoAdd(name, url string, opts *RepoAddOptions) error {
	if name == "" || url == "" {
		return errors.New("repo name and url are required")
	}

	args := []string{"repo", "add", name, url}
	if opts != nil {
		if opts.Username != "" && opts.Password != "" {
			args = append(args, "--username", opts.Username, "--password", opts.Password)
		}
		if opts.NoUpdate {
			args = append(args, "--no-update")
		}
	}

	_, err := w.exec(args...)
	return err
}

func (w Wrapper) RepoList() (repos []Repository, err error) {
	cmdArgs := []string{"repo", "list"}
	headerOutput := []string{"NAME", "URL"}
	captureRegex := `(?m)^(\S+)\s+(.*)$`
	err = w.listAction(cmdArgs, headerOutput, captureRegex, func(item []string) {
		repo := Repository{
			Name: item[1],
			URL:  item[2],
		}
		repos = append(repos, repo)
	})
	return
}

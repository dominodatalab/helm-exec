package helmexec

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

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
	out, err := w.exec("repo", "list")
	if err != nil {
		return
	}

	re := regexp.MustCompile(`(?m)^(\S+)\s+(.*)$`)
	matches := re.FindAllStringSubmatch(string(out), -1)
	header := []string{"NAME", "URL"}
	if matches == nil || !reflect.DeepEqual(strings.Fields(matches[0][0]), header) {
		return nil, fmt.Errorf("invalid list format: %s", string(out))
	}

	for _, data := range matches {
		if data[1] == "NAME" {
			continue
		}
		repo := Repository{
			Name: data[1],
			URL:  data[2],
		}
		repos = append(repos, repo)
	}
	return
}

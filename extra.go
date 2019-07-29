package helmexec

import "strings"

func (w Wrapper) IsRelease(name string) bool {
	out, err := w.exec("list", "--all", "--short", name)
	if err != nil {
		return false
	}

	for _, rlsName := range strings.Split(string(out), "\n") {
		if name == rlsName {
			return true
		}
	}
	return false
}

func (w Wrapper) IsRepo(name string) bool {
	repos, err := w.RepoList()
	if err != nil {
		return false
	}

	for _, r := range repos {
		if name == r.Name {
			return true
		}
	}
	return false
}

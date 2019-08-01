package helmexec

import (
	"encoding/json"
)

type release struct {
	Name string `json:"Name"`
}

type listResponse struct {
	Releases []release `json:"Releases"`
}

func (w Wrapper) IsRelease(name string) (exists bool) {
	out, err := w.exec("list", "--all", "--output=json", name)
	if err != nil {
		return
	}

	var rsp listResponse
	if err := json.Unmarshal(out, &rsp); err != nil {
		return
	}

	for _, rls := range rsp.Releases {
		if rls.Name == name {
			return true
		}
	}
	return
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

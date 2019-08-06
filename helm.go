package helmexec

type Helm interface {
	Init(...InitOption) error
	Install(chstr string, opts *InstallOptions) error
	//Upgrade(release string, ) error
	Delete(release string) error

	PluginInstall(pathOrURL, version string) error
	PluginList() ([]Plugin, error)

	RepoAdd(name, url string, opts *RepoAddOptions) error
	RepoList() ([]Repository, error)

	// convenience functions

	IsRepo(name string) bool
	IsRelease(name string) bool
}

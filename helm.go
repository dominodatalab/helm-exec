package helmexec

type Helm interface {
	Init(...InitOption) error

	//Install() error
	//Upgrade(release string, ) error
	//Delete(release string) error
	//
	//RepoAdd() error
	//
	//PluginInstall() error

	//IsRepo() bool
	//IsRelease() bool
}

package helmexec

type FakeRunner struct {
	execFn func(cmd []string) ([]byte, error)
}

func (m *FakeRunner) Execute(cmd string, args ...string) ([]byte, error) {
	fullCmd := append([]string{cmd}, args...)
	return m.execFn(fullCmd)
}

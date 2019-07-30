package helmexec

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/dominodatalab/vagrant-exec/command"
)

const binary = "helm"

type Wrapper struct {
	executable string
	runner     command.Runner
}

func New() Wrapper {
	return Wrapper{
		executable: binary,
		runner:     command.ShellRunner{},
	}
}

func (w *Wrapper) SetRunner(runner command.Runner) error {
	if runner != nil {
		w.runner = runner
		return nil
	}
	return errors.New("command runner cannot be nil")
}

func (w Wrapper) exec(args ...string) ([]byte, error) {
	return w.runner.Execute(w.executable, args...)
}

func (w Wrapper) listAction(cmdArgs, header []string, pattern string, fn func([]string)) error {
	bytes, err := w.exec(cmdArgs...)
	if err != nil {
		return err
	}
	output := string(bytes)

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(output, -1)
	if matches == nil || !reflect.DeepEqual(header, strings.Fields(matches[0][0])) {
		return fmt.Errorf("invalid list format, expected %q: %s", pattern, output)
	}
	for _, data := range matches[1:] {
		fn(data)
	}
	return nil
}

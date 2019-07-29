package helmexec

type initConfig struct {
	wait    bool
	upgrade bool
}

type InitOption func(*initConfig)

func InitWait(wait bool) InitOption {
	return func(opts *initConfig) {
		opts.wait = wait
	}
}

func InitUpgrade(upgrade bool) InitOption {
	return func(opts *initConfig) {
		opts.upgrade = upgrade
	}
}

func (w Wrapper) Init(opts ...InitOption) error {
	conf := new(initConfig)
	for _, opt := range opts {
		opt(conf)
	}

	args := []string{"init"}
	if conf.wait {
		args = append(args, "--wait")
	}
	if conf.upgrade {
		args = append(args, "--upgrade")
	}

	_, err := w.exec(args...)
	return err
}

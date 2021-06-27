package registry

type Registry struct {
	lgr Logger
}

func NewRegistry(l Logger) Registry {
	return Registry{
		lgr: l,
	}
}

func (r Registry) Logger() Logger {
	return r.lgr
}

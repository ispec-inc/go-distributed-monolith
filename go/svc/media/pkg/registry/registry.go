package registry

type Registry struct {
	repo Repository
	lgr  Logger
}

func NewRegistry(repo Repository, l Logger) Registry {
	return Registry{
		repo: repo,
		lgr:  l,
	}
}

func (r Registry) Logger() Logger {
	return r.lgr
}

func (r Registry) Repository() Repository {
	return r.repo
}

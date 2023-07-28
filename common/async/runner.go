package async

type Runner interface {
	Run() error
}

type RunFunc func() error

func (f RunFunc) Run() error {
	return f()
}

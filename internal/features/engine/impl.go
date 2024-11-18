package engine

type DefaultEngine struct {
	Config Config
}

func New(cfg Config) *DefaultEngine {
	return &DefaultEngine{
		Config: cfg,
	}
}

func (DefaultEngine) Handle() {

}

type Config struct {
}

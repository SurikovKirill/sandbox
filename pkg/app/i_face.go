package app

type Application interface {
	Run(name string, setup any)
}

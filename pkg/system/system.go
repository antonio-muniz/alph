package system

type System struct {
	componentProvider ComponentProvider
}

func New(componentProvider ComponentProvider) System {
	return System{componentProvider: componentProvider}
}

func (s System) Get(componentName string) interface{} {
	return s.componentProvider.Get(componentName)
}

func (s System) Clean() error {
	return s.componentProvider.Clean()
}

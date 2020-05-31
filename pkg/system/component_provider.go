package system

type ComponentProvider interface {
	Get(componentName string) interface{}
	Clean() error
}

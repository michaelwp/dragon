package dragon

type Router interface {
	GET(address string, handlerFunc HandlerFunc)
	POST(address string, handlerFunc HandlerFunc)
	Run(address string) error
	Group(address string) router
	Use(middleware ...func(*Dragon) error)
}

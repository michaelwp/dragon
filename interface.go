package dragon

type Router interface {
	GET(address string, handlerFunc HandlerFunc)
	POST(address string, handlerFunc HandlerFunc)
	PUT(address string, handlerFunc HandlerFunc)
	PATCH(address string, handlerFunc HandlerFunc)
	DELETE(address string, handlerFunc HandlerFunc)
	Run(address string) error
	Group(address string) router
	Use(middleware ...HandlerFunc)
}

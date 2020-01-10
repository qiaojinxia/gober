package gbinterface

type IRouter interface {
	PreHandle(request IRequest)
	Handler(request IRequest)
	PostHandle(request IRequest)
}

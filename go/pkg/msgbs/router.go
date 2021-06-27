package msgbs

type Router map[Event][]SubscribeFunc

func NewRouter() Router {
	return Router{}
}

func (r Router) Subscribe(evnt Event, subsc SubscribeFunc) {
	r[evnt] = append(r[evnt], subsc)
}

package events

type EventHandler func(event interface{})

type EventDispatcher struct {
	handlers map[string][]EventHandler
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (d *EventDispatcher) Register(eventName string, handler EventHandler) {
	d.handlers[eventName] = append(d.handlers[eventName], handler)
}

func (d *EventDispatcher) Dispatch(eventName string, event interface{}) {
	if handlers, ok := d.handlers[eventName]; ok {
		for _, handler := range handlers {
			handler(event)
		}
	}
}

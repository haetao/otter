package common

import "sync"

type HandlerFunc func(arg interface{}) interface{}

type FuncInfo struct {
	Tag  string
	Func HandlerFunc
}

type EventDispatcher struct {
	eventHandlers sync.Map
}

func (e *EventDispatcher) AddEventHandler(event string, fInfo FuncInfo) {
	handlers, ok := e.eventHandlers.Load(event)
	if !ok {
		handlers = make(map[string]FuncInfo, 0)
	}
	handlers2 := handlers.(map[string]FuncInfo)
	handlers2[fInfo.Tag] = fInfo
	e.eventHandlers.Store(event, handlers2)
}

func (e *EventDispatcher) RemoveEventHandler(event string, tag string) {
	handlers, ok := e.eventHandlers.Load(event)
	if !ok {
		return
	}
	handlers2 := handlers.(map[string]FuncInfo)
	delete(handlers2, tag)
}

func (e *EventDispatcher) Dispatch(event string, msg interface{}) interface{} {
	handlers, ok := e.eventHandlers.Load(event)
	if !ok {
		return nil
	}
	handlers2 := handlers.(map[string]FuncInfo)
	var result interface{}
	for _, info := range handlers2 {
		result = info.Func(msg)
	}
	if len(handlers2) == 1 {
		//只有一个handler的时候才有返回值
		return result
	}
	return nil
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		eventHandlers: sync.Map{},
	}
}

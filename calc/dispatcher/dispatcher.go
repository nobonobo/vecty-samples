package dispatcher

// Action ...
type Action interface{}

// ID ...
type ID int

var callbacks = []func(action Action){}

// Dispatch ...
func Dispatch(action Action) {
	for _, c := range callbacks {
		c(action)
	}
}

// Register ...
func Register(callback func(action Action)) ID {
	id := ID(len(callbacks))
	callbacks = append(callbacks, callback)
	return id
}

// Unregister ...
func Unregister(id ID) {
	callbacks = callbacks[:int(id)]
	remain := callbacks[int(id):]
	if len(remain) > 1 {
		callbacks = append(callbacks, remain[id+1:]...)
	}
}

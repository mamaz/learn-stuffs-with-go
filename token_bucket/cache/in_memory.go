package cache

type InMemory struct {
	internalMap map[string]interface{}
}

func New() *InMemory {
	inmemmory := &InMemory{
		internalMap: map[string]interface{}{},
	}

	return inmemmory
}

func (i *InMemory) Set(key string, value interface{}) {
	if i.internalMap == nil {
		i.internalMap = map[string]interface{}{}
	}

	i.internalMap[key] = value
}

func (i *InMemory) Get(key string) interface{} {
	value, ok := i.internalMap[key]
	if !ok {
		return nil
	}

	return value
}

func (i *InMemory) Delete(key string) {
	delete(i.internalMap, key)
}

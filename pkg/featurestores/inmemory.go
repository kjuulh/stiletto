package featurestores

type InMemory struct {
	store []string
}

func NewInMemoryFeatureStore(data []string) FeatureStore {
	return &InMemory{
		store: data,
	}
}

func (i InMemory) Get(key string) (bool, error) {
	for _, dataKey := range i.store {
		if dataKey == key {
			return true, nil
		}
	}

	return false, nil
}

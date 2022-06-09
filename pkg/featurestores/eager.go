package featurestores

import "fmt"

type Eager struct {
	inner FeatureStore
}

func NewEagerFeatureStore(f FeatureStore, eagerFetchedKeys []string) FeatureStore {
	for _, k := range eagerFetchedKeys {
		_, err := f.Get(k)
		if err != nil {
			fmt.Printf("eager load of key failed for key: %s", k)
		}
	}

	return &Eager{
		inner: f,
	}
}

func (i *Eager) Get(key string) (bool, error) {
	return i.inner.Get(key)
}

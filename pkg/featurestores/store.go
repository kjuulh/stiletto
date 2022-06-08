package featurestores

type FeatureStore interface {
	Get(key string) (bool, error)
}

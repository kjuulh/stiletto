package featurestores

//go:generate moq -out store_moq.go . FeatureStore

type FeatureStore interface {
	Get(key string) (bool, error)
}

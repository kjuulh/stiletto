package client

import (
	"context"
	"fmt"
	"git.front.kjuulh.io/stiletto/stiletto/pkg/featurestores"
)

type StilettoClient interface {
	SetFeatureStore(featureKey string, store featurestores.FeatureStore)
	GetCapture(ctx context.Context, featureKey string, queryKey string) (bool, error)
}

type client struct {
	features map[string]featurestores.FeatureStore
}

func NewStilettoClient() StilettoClient {
	return &client{
		features: make(map[string]featurestores.FeatureStore),
	}
}

func (c client) SetFeatureStore(featureKey string, store featurestores.FeatureStore) {
	c.features[featureKey] = store
}

func (c client) GetCapture(ctx context.Context, featureKey string, queryKey string) (bool, error) {
	feature, ok := c.features[featureKey]
	if !ok {
		return false, fmt.Errorf("feature not found. feature: %s", featureKey)
	}

	capture, err := feature.Get(queryKey)
	if err != nil {
		return false, fmt.Errorf("could not get entry, %w", err)
	}

	return capture, nil
}

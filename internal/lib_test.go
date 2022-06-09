package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kjuulh/stiletto/pkg/client"
	"github.com/kjuulh/stiletto/pkg/featurestores"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIsEntry(t *testing.T) {
	uut := client.NewStilettoClient()

	userId := "some-userID"
	featureStore := featurestores.NewInMemoryFeatureStore([]string{
		userId,
	})
	uut.SetFeatureStore("user", featureStore)

	isEntry, err := uut.GetFeature(context.Background(), "user", userId)
	require.NoError(t, err)
	require.True(t, isEntry)

	isNotEntry, err := uut.GetFeature(context.Background(), "user", "not a correct user id")
	require.NoError(t, err)
	require.False(t, isNotEntry)
}

func TestStack(t *testing.T) {
	uut := client.NewStilettoClient()

	httpClient := &featurestores.HttpClientMock{GetFunc: func(s string) ([]byte, error) {
		if s == "http://localhost/keys/some-key" {
			return json.Marshal(featurestores.Feature{
				Key:     "some-key",
				Enabled: true,
			})
		}
		if s == "http://localhost/keys/some-other-key" {
			return json.Marshal(featurestores.Feature{
				Key:     "some-other-key",
				Enabled: false,
			})
		}
		return nil, fmt.Errorf("query does not match aborting, %s", s)
	}}

	uut.SetFeatureStore(
		"user",
		featurestores.NewEagerFeatureStore(
			featurestores.NewCache(
				featurestores.
					NewRemote("http://localhost/keys").
					WithHttpClient(httpClient).
					Build()),
			[]string{"some-key", "some-other-key", "some-key"},
		),
	)

	feature, err := uut.GetFeature(context.Background(), "user", "some-key")
	require.NoError(t, err)
	require.True(t, feature)

	require.Equal(t, 2, len(httpClient.GetCalls()))
}

package internal

import (
	"context"
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

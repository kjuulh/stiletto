package featurestores

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewEager(t *testing.T) {
	type testData struct {
		Name         string
		FeatureStore *FeatureStoreMock
		Keys         []string
		Expected     struct {
			Get []struct {
				Key string
			}
		}
	}

	tt := []testData{
		{
			Name: "empty set of keys",
			FeatureStore: &FeatureStoreMock{
				GetFunc: func(key string) (bool, error) {
					return false, nil
				},
			},
			Keys:     []string{},
			Expected: struct{ Get []struct{ Key string } }{},
		},
		{
			Name: "single key",
			FeatureStore: &FeatureStoreMock{
				GetFunc: func(key string) (bool, error) {
					return false, nil
				},
			},
			Keys: []string{"some-single-key"},
			Expected: struct{ Get []struct{ Key string } }{
				Get: []struct{ Key string }{{Key: "some-single-key"}}},
		},
		{
			Name: "two keys",
			FeatureStore: &FeatureStoreMock{
				GetFunc: func(key string) (bool, error) {
					return false, nil
				},
			},
			Keys: []string{"some-single-key", "some-second-key"},
			Expected: struct{ Get []struct{ Key string } }{
				Get: []struct{ Key string }{
					{Key: "some-single-key"},
					{Key: "some-second-key"},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {

			uut := NewEagerFeatureStore(tc.FeatureStore, tc.Keys)

			require.NotNil(t, uut)
			require.Equal(t, tc.Expected, tc.FeatureStore.calls)
		})
	}
}

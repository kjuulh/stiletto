package featurestores

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCacheGet(t *testing.T) {
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
			Name: "cache miss on single key",
			FeatureStore: &FeatureStoreMock{
				GetFunc: func(key string) (bool, error) {
					return false, nil
				},
			},
			Keys: []string{"first-key-will-miss"},
			Expected: struct{ Get []struct{ Key string } }{
				Get: []struct{ Key string }{
					{Key: "first-key-will-miss"},
				},
			},
		},
		{
			Name: "cache hit on second hit",
			FeatureStore: &FeatureStoreMock{
				GetFunc: func(key string) (bool, error) {
					return false, nil
				},
			},
			Keys: []string{"second-key-will-hit", "second-key-will-hit"},
			Expected: struct{ Get []struct{ Key string } }{
				Get: []struct{ Key string }{
					{Key: "second-key-will-hit"},
				},
			},
		},
		{
			Name: "mix will still hit on second",
			FeatureStore: &FeatureStoreMock{
				GetFunc: func(key string) (bool, error) {
					return false, nil
				},
			},
			Keys: []string{"second-key-will-hit", "mixed-key-in-here", "second-key-will-hit"},
			Expected: struct{ Get []struct{ Key string } }{
				Get: []struct{ Key string }{
					{Key: "second-key-will-hit"},
					{Key: "mixed-key-in-here"},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			uut := NewCache(tc.FeatureStore)

			for _, k := range tc.Keys {
				res, err := uut.Get(k)
				require.NoError(t, err)
				require.False(t, res)
			}

			require.Equal(t, tc.Expected, tc.FeatureStore.calls)
		})
	}
}

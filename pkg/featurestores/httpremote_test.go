package featurestores

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHttpRemoteGet(t *testing.T) {
	type testData struct {
		Name                  string
		HttpClient            *HttpClientMock
		Key                   string
		ExpectedAmountOfCalls int
		ExpectedFeatureValue  bool
	}

	tt := []testData{
		{
			Name: "matched url returns feature",
			HttpClient: &HttpClientMock{
				GetFunc: func(query string) ([]byte, error) {
					if query == "http://test-url-does-not-exist-during-unit-test/test-key" {
						return json.Marshal(Feature{
							Key:     "test-key",
							Enabled: true,
						})
					}
					return nil, fmt.Errorf("query does not match call")
				},
			},
			Key:                   "test-key",
			ExpectedAmountOfCalls: 1,
			ExpectedFeatureValue:  true,
		},
		{
			Name: "matched url returns feature alternative",
			HttpClient: &HttpClientMock{
				GetFunc: func(query string) ([]byte, error) {
					if query == "http://test-url-does-not-exist-during-unit-test/test-key-some-other" {
						return json.Marshal(Feature{
							Key:     "test-key-some-other",
							Enabled: false,
						})
					}
					return nil, fmt.Errorf("query does not match call")
				},
			},
			Key:                   "test-key-some-other",
			ExpectedAmountOfCalls: 1,
			ExpectedFeatureValue:  false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			uut := NewRemote("http://test-url-does-not-exist-during-unit-test").
				WithHttpClient(tc.HttpClient).
				Build()

			actual, err := uut.Get(tc.Key)

			require.NoError(t, err)
			require.Equal(t, tc.ExpectedAmountOfCalls, len(tc.HttpClient.GetCalls()))
			require.Equal(t, tc.ExpectedFeatureValue, actual)
		})
	}
}

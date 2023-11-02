/*
Cashfree Payment Gateway APIs

Testing RefundsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cashfree_pg

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	cashfree "github.com/cashfree/cashfree-pg/v3"
)

func Test_cashfree_pg_RefundsAPIService(t *testing.T) {

	configuration := cashfree.NewConfiguration()
	apiClient := cashfree.NewAPIClient(configuration)

	t.Run("Test RefundsAPIService PGOrderCreateRefund", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId string

		resp, httpRes, err := apiClient.RefundsAPI.PGOrderCreateRefund(context.Background(), orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RefundsAPIService PGOrderFetchRefund", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId string
		var refundId string

		resp, httpRes, err := apiClient.RefundsAPI.PGOrderFetchRefund(context.Background(), orderId, refundId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RefundsAPIService PGOrderFetchRefunds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orderId string

		resp, httpRes, err := apiClient.RefundsAPI.PGOrderFetchRefunds(context.Background(), orderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

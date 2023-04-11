package pricings

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSubscriptionPricingOperationResponse struct {
	HttpResponse *http.Response
	Model        *Pricing
}

// GetSubscriptionPricing ...
func (c PricingsClient) GetSubscriptionPricing(ctx context.Context, id PricingId) (result GetSubscriptionPricingOperationResponse, err error) {
	req, err := c.preparerForGetSubscriptionPricing(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricings.PricingsClient", "GetSubscriptionPricing", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricings.PricingsClient", "GetSubscriptionPricing", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSubscriptionPricing(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricings.PricingsClient", "GetSubscriptionPricing", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSubscriptionPricing prepares the GetSubscriptionPricing request.
func (c PricingsClient) preparerForGetSubscriptionPricing(ctx context.Context, id PricingId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetSubscriptionPricing handles the response to the GetSubscriptionPricing request. The method always
// closes the http.Response Body.
func (c PricingsClient) responderForGetSubscriptionPricing(resp *http.Response) (result GetSubscriptionPricingOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

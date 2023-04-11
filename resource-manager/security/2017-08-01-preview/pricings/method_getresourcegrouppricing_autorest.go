package pricings

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetResourceGroupPricingOperationResponse struct {
	HttpResponse *http.Response
	Model        *Pricing
}

// GetResourceGroupPricing ...
func (c PricingsClient) GetResourceGroupPricing(ctx context.Context, id ProviderPricingId) (result GetResourceGroupPricingOperationResponse, err error) {
	req, err := c.preparerForGetResourceGroupPricing(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricings.PricingsClient", "GetResourceGroupPricing", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricings.PricingsClient", "GetResourceGroupPricing", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetResourceGroupPricing(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricings.PricingsClient", "GetResourceGroupPricing", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetResourceGroupPricing prepares the GetResourceGroupPricing request.
func (c PricingsClient) preparerForGetResourceGroupPricing(ctx context.Context, id ProviderPricingId) (*http.Request, error) {
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

// responderForGetResourceGroupPricing handles the response to the GetResourceGroupPricing request. The method always
// closes the http.Response Body.
func (c PricingsClient) responderForGetResourceGroupPricing(resp *http.Response) (result GetResourceGroupPricingOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

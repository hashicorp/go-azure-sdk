package pricings

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateResourceGroupPricingOperationResponse struct {
	HttpResponse *http.Response
	Model        *Pricing
}

// CreateOrUpdateResourceGroupPricing ...
func (c PricingsClient) CreateOrUpdateResourceGroupPricing(ctx context.Context, id ProviderPricingId, input Pricing) (result CreateOrUpdateResourceGroupPricingOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateResourceGroupPricing(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricings.PricingsClient", "CreateOrUpdateResourceGroupPricing", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricings.PricingsClient", "CreateOrUpdateResourceGroupPricing", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateResourceGroupPricing(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricings.PricingsClient", "CreateOrUpdateResourceGroupPricing", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateResourceGroupPricing prepares the CreateOrUpdateResourceGroupPricing request.
func (c PricingsClient) preparerForCreateOrUpdateResourceGroupPricing(ctx context.Context, id ProviderPricingId, input Pricing) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateResourceGroupPricing handles the response to the CreateOrUpdateResourceGroupPricing request. The method always
// closes the http.Response Body.
func (c PricingsClient) responderForCreateOrUpdateResourceGroupPricing(resp *http.Response) (result CreateOrUpdateResourceGroupPricingOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

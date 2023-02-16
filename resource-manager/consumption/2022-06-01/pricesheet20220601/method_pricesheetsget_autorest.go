package pricesheet20220601

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PriceSheetsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *PriceSheetResultV2
}

type PriceSheetsGetOperationOptions struct {
	Expand *string
	Top    *int64
}

func DefaultPriceSheetsGetOperationOptions() PriceSheetsGetOperationOptions {
	return PriceSheetsGetOperationOptions{}
}

func (o PriceSheetsGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o PriceSheetsGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// PriceSheetsGet ...
func (c PriceSheet20220601Client) PriceSheetsGet(ctx context.Context, id commonids.SubscriptionId, options PriceSheetsGetOperationOptions) (result PriceSheetsGetOperationResponse, err error) {
	req, err := c.preparerForPriceSheetsGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricesheet20220601.PriceSheet20220601Client", "PriceSheetsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricesheet20220601.PriceSheet20220601Client", "PriceSheetsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPriceSheetsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricesheet20220601.PriceSheet20220601Client", "PriceSheetsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPriceSheetsGet prepares the PriceSheetsGet request.
func (c PriceSheet20220601Client) preparerForPriceSheetsGet(ctx context.Context, id commonids.SubscriptionId, options PriceSheetsGetOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Consumption/pricesheets/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPriceSheetsGet handles the response to the PriceSheetsGet request. The method always
// closes the http.Response Body.
func (c PriceSheet20220601Client) responderForPriceSheetsGet(resp *http.Response) (result PriceSheetsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package pricesheet

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetByBillingPeriodOperationResponse struct {
	HttpResponse *http.Response
	Model        *PriceSheetResult
}

type GetByBillingPeriodOperationOptions struct {
	Expand *string
	Top    *int64
}

func DefaultGetByBillingPeriodOperationOptions() GetByBillingPeriodOperationOptions {
	return GetByBillingPeriodOperationOptions{}
}

func (o GetByBillingPeriodOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetByBillingPeriodOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// GetByBillingPeriod ...
func (c PriceSheetClient) GetByBillingPeriod(ctx context.Context, id BillingPeriodId, options GetByBillingPeriodOperationOptions) (result GetByBillingPeriodOperationResponse, err error) {
	req, err := c.preparerForGetByBillingPeriod(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricesheet.PriceSheetClient", "GetByBillingPeriod", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricesheet.PriceSheetClient", "GetByBillingPeriod", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetByBillingPeriod(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricesheet.PriceSheetClient", "GetByBillingPeriod", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetByBillingPeriod prepares the GetByBillingPeriod request.
func (c PriceSheetClient) preparerForGetByBillingPeriod(ctx context.Context, id BillingPeriodId, options GetByBillingPeriodOperationOptions) (*http.Request, error) {
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

// responderForGetByBillingPeriod handles the response to the GetByBillingPeriod request. The method always
// closes the http.Response Body.
func (c PriceSheetClient) responderForGetByBillingPeriod(resp *http.Response) (result GetByBillingPeriodOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

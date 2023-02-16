package pricesheet20220601

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PriceSheetsGetByBillingPeriodOperationResponse struct {
	HttpResponse *http.Response
	Model        *PriceSheetResultV2
}

type PriceSheetsGetByBillingPeriodOperationOptions struct {
	Expand *string
	Top    *int64
}

func DefaultPriceSheetsGetByBillingPeriodOperationOptions() PriceSheetsGetByBillingPeriodOperationOptions {
	return PriceSheetsGetByBillingPeriodOperationOptions{}
}

func (o PriceSheetsGetByBillingPeriodOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o PriceSheetsGetByBillingPeriodOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// PriceSheetsGetByBillingPeriod ...
func (c PriceSheet20220601Client) PriceSheetsGetByBillingPeriod(ctx context.Context, id BillingPeriodId, options PriceSheetsGetByBillingPeriodOperationOptions) (result PriceSheetsGetByBillingPeriodOperationResponse, err error) {
	req, err := c.preparerForPriceSheetsGetByBillingPeriod(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricesheet20220601.PriceSheet20220601Client", "PriceSheetsGetByBillingPeriod", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricesheet20220601.PriceSheet20220601Client", "PriceSheetsGetByBillingPeriod", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPriceSheetsGetByBillingPeriod(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "pricesheet20220601.PriceSheet20220601Client", "PriceSheetsGetByBillingPeriod", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPriceSheetsGetByBillingPeriod prepares the PriceSheetsGetByBillingPeriod request.
func (c PriceSheet20220601Client) preparerForPriceSheetsGetByBillingPeriod(ctx context.Context, id BillingPeriodId, options PriceSheetsGetByBillingPeriodOperationOptions) (*http.Request, error) {
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

// responderForPriceSheetsGetByBillingPeriod handles the response to the PriceSheetsGetByBillingPeriod request. The method always
// closes the http.Response Body.
func (c PriceSheet20220601Client) responderForPriceSheetsGetByBillingPeriod(resp *http.Response) (result PriceSheetsGetByBillingPeriodOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

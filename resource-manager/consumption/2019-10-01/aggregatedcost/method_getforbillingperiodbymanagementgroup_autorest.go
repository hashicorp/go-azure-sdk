package aggregatedcost

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetForBillingPeriodByManagementGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *ManagementGroupAggregatedCostResult
}

// GetForBillingPeriodByManagementGroup ...
func (c AggregatedCostClient) GetForBillingPeriodByManagementGroup(ctx context.Context, id Providers2BillingPeriodId) (result GetForBillingPeriodByManagementGroupOperationResponse, err error) {
	req, err := c.preparerForGetForBillingPeriodByManagementGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aggregatedcost.AggregatedCostClient", "GetForBillingPeriodByManagementGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "aggregatedcost.AggregatedCostClient", "GetForBillingPeriodByManagementGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetForBillingPeriodByManagementGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aggregatedcost.AggregatedCostClient", "GetForBillingPeriodByManagementGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetForBillingPeriodByManagementGroup prepares the GetForBillingPeriodByManagementGroup request.
func (c AggregatedCostClient) preparerForGetForBillingPeriodByManagementGroup(ctx context.Context, id Providers2BillingPeriodId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/Microsoft.Consumption/aggregatedCost", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetForBillingPeriodByManagementGroup handles the response to the GetForBillingPeriodByManagementGroup request. The method always
// closes the http.Response Body.
func (c AggregatedCostClient) responderForGetForBillingPeriodByManagementGroup(resp *http.Response) (result GetForBillingPeriodByManagementGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

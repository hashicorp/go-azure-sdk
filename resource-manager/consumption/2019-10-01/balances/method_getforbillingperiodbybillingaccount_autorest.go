package balances

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetForBillingPeriodByBillingAccountOperationResponse struct {
	HttpResponse *http.Response
	Model        *Balance
}

// GetForBillingPeriodByBillingAccount ...
func (c BalancesClient) GetForBillingPeriodByBillingAccount(ctx context.Context, id BillingAccountBillingPeriodId) (result GetForBillingPeriodByBillingAccountOperationResponse, err error) {
	req, err := c.preparerForGetForBillingPeriodByBillingAccount(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "balances.BalancesClient", "GetForBillingPeriodByBillingAccount", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "balances.BalancesClient", "GetForBillingPeriodByBillingAccount", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetForBillingPeriodByBillingAccount(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "balances.BalancesClient", "GetForBillingPeriodByBillingAccount", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetForBillingPeriodByBillingAccount prepares the GetForBillingPeriodByBillingAccount request.
func (c BalancesClient) preparerForGetForBillingPeriodByBillingAccount(ctx context.Context, id BillingAccountBillingPeriodId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Consumption/balances", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetForBillingPeriodByBillingAccount handles the response to the GetForBillingPeriodByBillingAccount request. The method always
// closes the http.Response Body.
func (c BalancesClient) responderForGetForBillingPeriodByBillingAccount(resp *http.Response) (result GetForBillingPeriodByBillingAccountOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

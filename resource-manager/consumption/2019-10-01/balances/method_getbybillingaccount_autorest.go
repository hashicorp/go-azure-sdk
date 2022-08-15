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

type GetByBillingAccountOperationResponse struct {
	HttpResponse *http.Response
	Model        *Balance
}

// GetByBillingAccount ...
func (c BalancesClient) GetByBillingAccount(ctx context.Context, id BillingAccountId) (result GetByBillingAccountOperationResponse, err error) {
	req, err := c.preparerForGetByBillingAccount(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "balances.BalancesClient", "GetByBillingAccount", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "balances.BalancesClient", "GetByBillingAccount", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetByBillingAccount(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "balances.BalancesClient", "GetByBillingAccount", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetByBillingAccount prepares the GetByBillingAccount request.
func (c BalancesClient) preparerForGetByBillingAccount(ctx context.Context, id BillingAccountId) (*http.Request, error) {
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

// responderForGetByBillingAccount handles the response to the GetByBillingAccount request. The method always
// closes the http.Response Body.
func (c BalancesClient) responderForGetByBillingAccount(resp *http.Response) (result GetByBillingAccountOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

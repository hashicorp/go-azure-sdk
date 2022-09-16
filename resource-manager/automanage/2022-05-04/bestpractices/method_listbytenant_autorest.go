package bestpractices

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByTenantOperationResponse struct {
	HttpResponse *http.Response
	Model        *BestPracticeList
}

// ListByTenant ...
func (c BestPracticesClient) ListByTenant(ctx context.Context) (result ListByTenantOperationResponse, err error) {
	req, err := c.preparerForListByTenant(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bestpractices.BestPracticesClient", "ListByTenant", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bestpractices.BestPracticesClient", "ListByTenant", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByTenant(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bestpractices.BestPracticesClient", "ListByTenant", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByTenant prepares the ListByTenant request.
func (c BestPracticesClient) preparerForListByTenant(ctx context.Context) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath("/providers/Microsoft.AutoManage/bestPractices"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByTenant handles the response to the ListByTenant request. The method always
// closes the http.Response Body.
func (c BestPracticesClient) responderForListByTenant(resp *http.Response) (result ListByTenantOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

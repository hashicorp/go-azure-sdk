package usages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByAutomationAccountOperationResponse struct {
	HttpResponse *http.Response
	Model        *UsageListResult
}

// ListByAutomationAccount ...
func (c UsagesClient) ListByAutomationAccount(ctx context.Context, id AutomationAccountId) (result ListByAutomationAccountOperationResponse, err error) {
	req, err := c.preparerForListByAutomationAccount(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "usages.UsagesClient", "ListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "usages.UsagesClient", "ListByAutomationAccount", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByAutomationAccount(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "usages.UsagesClient", "ListByAutomationAccount", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByAutomationAccount prepares the ListByAutomationAccount request.
func (c UsagesClient) preparerForListByAutomationAccount(ctx context.Context, id AutomationAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/usages", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByAutomationAccount handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (c UsagesClient) responderForListByAutomationAccount(resp *http.Response) (result ListByAutomationAccountOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

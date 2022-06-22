package statistics

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
	Model        *StatisticsListResult
}

type ListByAutomationAccountOperationOptions struct {
	Filter *string
}

func DefaultListByAutomationAccountOperationOptions() ListByAutomationAccountOperationOptions {
	return ListByAutomationAccountOperationOptions{}
}

func (o ListByAutomationAccountOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByAutomationAccountOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ListByAutomationAccount ...
func (c StatisticsClient) ListByAutomationAccount(ctx context.Context, id AutomationAccountId, options ListByAutomationAccountOperationOptions) (result ListByAutomationAccountOperationResponse, err error) {
	req, err := c.preparerForListByAutomationAccount(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "statistics.StatisticsClient", "ListByAutomationAccount", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "statistics.StatisticsClient", "ListByAutomationAccount", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByAutomationAccount(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "statistics.StatisticsClient", "ListByAutomationAccount", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByAutomationAccount prepares the ListByAutomationAccount request.
func (c StatisticsClient) preparerForListByAutomationAccount(ctx context.Context, id AutomationAccountId, options ListByAutomationAccountOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/statistics", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByAutomationAccount handles the response to the ListByAutomationAccount request. The method always
// closes the http.Response Body.
func (c StatisticsClient) responderForListByAutomationAccount(resp *http.Response) (result ListByAutomationAccountOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

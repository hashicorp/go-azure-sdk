package reports

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByRequestOperationResponse struct {
	HttpResponse *http.Response
	Model        *RequestReportCollection
}

type ListByRequestOperationOptions struct {
	Filter *string
	Skip   *int64
	Top    *int64
}

func DefaultListByRequestOperationOptions() ListByRequestOperationOptions {
	return ListByRequestOperationOptions{}
}

func (o ListByRequestOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListByRequestOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListByRequest ...
func (c ReportsClient) ListByRequest(ctx context.Context, id ServiceId, options ListByRequestOperationOptions) (result ListByRequestOperationResponse, err error) {
	req, err := c.preparerForListByRequest(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByRequest", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByRequest", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByRequest(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reports.ReportsClient", "ListByRequest", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByRequest prepares the ListByRequest request.
func (c ReportsClient) preparerForListByRequest(ctx context.Context, id ServiceId, options ListByRequestOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/reports/byRequest", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByRequest handles the response to the ListByRequest request. The method always
// closes the http.Response Body.
func (c ReportsClient) responderForListByRequest(resp *http.Response) (result ListByRequestOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package restorables

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableGremlinGraphsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorableGremlinGraphsListResult
}

type RestorableGremlinGraphsListOperationOptions struct {
	EndTime                      *string
	RestorableGremlinDatabaseRid *string
	StartTime                    *string
}

func DefaultRestorableGremlinGraphsListOperationOptions() RestorableGremlinGraphsListOperationOptions {
	return RestorableGremlinGraphsListOperationOptions{}
}

func (o RestorableGremlinGraphsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RestorableGremlinGraphsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.EndTime != nil {
		out["endTime"] = *o.EndTime
	}

	if o.RestorableGremlinDatabaseRid != nil {
		out["restorableGremlinDatabaseRid"] = *o.RestorableGremlinDatabaseRid
	}

	if o.StartTime != nil {
		out["startTime"] = *o.StartTime
	}

	return out
}

// RestorableGremlinGraphsList ...
func (c RestorablesClient) RestorableGremlinGraphsList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableGremlinGraphsListOperationOptions) (result RestorableGremlinGraphsListOperationResponse, err error) {
	req, err := c.preparerForRestorableGremlinGraphsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableGremlinGraphsList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableGremlinGraphsList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestorableGremlinGraphsList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableGremlinGraphsList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestorableGremlinGraphsList prepares the RestorableGremlinGraphsList request.
func (c RestorablesClient) preparerForRestorableGremlinGraphsList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableGremlinGraphsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/restorableGraphs", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestorableGremlinGraphsList handles the response to the RestorableGremlinGraphsList request. The method always
// closes the http.Response Body.
func (c RestorablesClient) responderForRestorableGremlinGraphsList(resp *http.Response) (result RestorableGremlinGraphsListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

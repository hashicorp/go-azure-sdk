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

type RestorableGremlinResourcesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorableGremlinResourcesListResult
}

type RestorableGremlinResourcesListOperationOptions struct {
	RestoreLocation       *string
	RestoreTimestampInUtc *string
}

func DefaultRestorableGremlinResourcesListOperationOptions() RestorableGremlinResourcesListOperationOptions {
	return RestorableGremlinResourcesListOperationOptions{}
}

func (o RestorableGremlinResourcesListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RestorableGremlinResourcesListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.RestoreLocation != nil {
		out["restoreLocation"] = *o.RestoreLocation
	}

	if o.RestoreTimestampInUtc != nil {
		out["restoreTimestampInUtc"] = *o.RestoreTimestampInUtc
	}

	return out
}

// RestorableGremlinResourcesList ...
func (c RestorablesClient) RestorableGremlinResourcesList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableGremlinResourcesListOperationOptions) (result RestorableGremlinResourcesListOperationResponse, err error) {
	req, err := c.preparerForRestorableGremlinResourcesList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableGremlinResourcesList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableGremlinResourcesList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestorableGremlinResourcesList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableGremlinResourcesList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestorableGremlinResourcesList prepares the RestorableGremlinResourcesList request.
func (c RestorablesClient) preparerForRestorableGremlinResourcesList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableGremlinResourcesListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/restorableGremlinResources", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestorableGremlinResourcesList handles the response to the RestorableGremlinResourcesList request. The method always
// closes the http.Response Body.
func (c RestorablesClient) responderForRestorableGremlinResourcesList(resp *http.Response) (result RestorableGremlinResourcesListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

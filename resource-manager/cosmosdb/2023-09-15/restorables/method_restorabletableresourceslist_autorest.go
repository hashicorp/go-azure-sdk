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

type RestorableTableResourcesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorableTableResourcesListResult
}

type RestorableTableResourcesListOperationOptions struct {
	RestoreLocation       *string
	RestoreTimestampInUtc *string
}

func DefaultRestorableTableResourcesListOperationOptions() RestorableTableResourcesListOperationOptions {
	return RestorableTableResourcesListOperationOptions{}
}

func (o RestorableTableResourcesListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RestorableTableResourcesListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.RestoreLocation != nil {
		out["restoreLocation"] = *o.RestoreLocation
	}

	if o.RestoreTimestampInUtc != nil {
		out["restoreTimestampInUtc"] = *o.RestoreTimestampInUtc
	}

	return out
}

// RestorableTableResourcesList ...
func (c RestorablesClient) RestorableTableResourcesList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableTableResourcesListOperationOptions) (result RestorableTableResourcesListOperationResponse, err error) {
	req, err := c.preparerForRestorableTableResourcesList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableTableResourcesList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableTableResourcesList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestorableTableResourcesList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableTableResourcesList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestorableTableResourcesList prepares the RestorableTableResourcesList request.
func (c RestorablesClient) preparerForRestorableTableResourcesList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableTableResourcesListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/restorableTableResources", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestorableTableResourcesList handles the response to the RestorableTableResourcesList request. The method always
// closes the http.Response Body.
func (c RestorablesClient) responderForRestorableTableResourcesList(resp *http.Response) (result RestorableTableResourcesListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

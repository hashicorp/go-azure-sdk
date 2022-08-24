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

type RestorableSqlResourcesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorableSqlResourcesListResult
}

type RestorableSqlResourcesListOperationOptions struct {
	RestoreLocation       *string
	RestoreTimestampInUtc *string
}

func DefaultRestorableSqlResourcesListOperationOptions() RestorableSqlResourcesListOperationOptions {
	return RestorableSqlResourcesListOperationOptions{}
}

func (o RestorableSqlResourcesListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RestorableSqlResourcesListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.RestoreLocation != nil {
		out["restoreLocation"] = *o.RestoreLocation
	}

	if o.RestoreTimestampInUtc != nil {
		out["restoreTimestampInUtc"] = *o.RestoreTimestampInUtc
	}

	return out
}

// RestorableSqlResourcesList ...
func (c RestorablesClient) RestorableSqlResourcesList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableSqlResourcesListOperationOptions) (result RestorableSqlResourcesListOperationResponse, err error) {
	req, err := c.preparerForRestorableSqlResourcesList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableSqlResourcesList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableSqlResourcesList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestorableSqlResourcesList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableSqlResourcesList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestorableSqlResourcesList prepares the RestorableSqlResourcesList request.
func (c RestorablesClient) preparerForRestorableSqlResourcesList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableSqlResourcesListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/restorableSqlResources", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestorableSqlResourcesList handles the response to the RestorableSqlResourcesList request. The method always
// closes the http.Response Body.
func (c RestorablesClient) responderForRestorableSqlResourcesList(resp *http.Response) (result RestorableSqlResourcesListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

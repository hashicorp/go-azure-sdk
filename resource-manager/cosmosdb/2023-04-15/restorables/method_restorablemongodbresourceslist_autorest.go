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

type RestorableMongodbResourcesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorableMongodbResourcesListResult
}

type RestorableMongodbResourcesListOperationOptions struct {
	RestoreLocation       *string
	RestoreTimestampInUtc *string
}

func DefaultRestorableMongodbResourcesListOperationOptions() RestorableMongodbResourcesListOperationOptions {
	return RestorableMongodbResourcesListOperationOptions{}
}

func (o RestorableMongodbResourcesListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RestorableMongodbResourcesListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.RestoreLocation != nil {
		out["restoreLocation"] = *o.RestoreLocation
	}

	if o.RestoreTimestampInUtc != nil {
		out["restoreTimestampInUtc"] = *o.RestoreTimestampInUtc
	}

	return out
}

// RestorableMongodbResourcesList ...
func (c RestorablesClient) RestorableMongodbResourcesList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableMongodbResourcesListOperationOptions) (result RestorableMongodbResourcesListOperationResponse, err error) {
	req, err := c.preparerForRestorableMongodbResourcesList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableMongodbResourcesList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableMongodbResourcesList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestorableMongodbResourcesList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableMongodbResourcesList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestorableMongodbResourcesList prepares the RestorableMongodbResourcesList request.
func (c RestorablesClient) preparerForRestorableMongodbResourcesList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableMongodbResourcesListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/restorableMongodbResources", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestorableMongodbResourcesList handles the response to the RestorableMongodbResourcesList request. The method always
// closes the http.Response Body.
func (c RestorablesClient) responderForRestorableMongodbResourcesList(resp *http.Response) (result RestorableMongodbResourcesListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

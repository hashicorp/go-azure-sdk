package compute

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorePointsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorePoint
}

type RestorePointsGetOperationOptions struct {
	Expand *RestorePointExpandOptions
}

func DefaultRestorePointsGetOperationOptions() RestorePointsGetOperationOptions {
	return RestorePointsGetOperationOptions{}
}

func (o RestorePointsGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RestorePointsGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	return out
}

// RestorePointsGet ...
func (c ComputeClient) RestorePointsGet(ctx context.Context, id RestorePointId, options RestorePointsGetOperationOptions) (result RestorePointsGetOperationResponse, err error) {
	req, err := c.preparerForRestorePointsGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ComputeClient", "RestorePointsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ComputeClient", "RestorePointsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestorePointsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.ComputeClient", "RestorePointsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestorePointsGet prepares the RestorePointsGet request.
func (c ComputeClient) preparerForRestorePointsGet(ctx context.Context, id RestorePointId, options RestorePointsGetOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestorePointsGet handles the response to the RestorePointsGet request. The method always
// closes the http.Response Body.
func (c ComputeClient) responderForRestorePointsGet(resp *http.Response) (result RestorePointsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

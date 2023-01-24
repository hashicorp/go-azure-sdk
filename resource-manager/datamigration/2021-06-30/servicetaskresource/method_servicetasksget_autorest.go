package servicetaskresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceTasksGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProjectTask
}

type ServiceTasksGetOperationOptions struct {
	Expand *string
}

func DefaultServiceTasksGetOperationOptions() ServiceTasksGetOperationOptions {
	return ServiceTasksGetOperationOptions{}
}

func (o ServiceTasksGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ServiceTasksGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	return out
}

// ServiceTasksGet ...
func (c ServiceTaskResourceClient) ServiceTasksGet(ctx context.Context, id ServiceTaskId, options ServiceTasksGetOperationOptions) (result ServiceTasksGetOperationResponse, err error) {
	req, err := c.preparerForServiceTasksGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicetaskresource.ServiceTaskResourceClient", "ServiceTasksGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicetaskresource.ServiceTaskResourceClient", "ServiceTasksGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForServiceTasksGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicetaskresource.ServiceTaskResourceClient", "ServiceTasksGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForServiceTasksGet prepares the ServiceTasksGet request.
func (c ServiceTaskResourceClient) preparerForServiceTasksGet(ctx context.Context, id ServiceTaskId, options ServiceTasksGetOperationOptions) (*http.Request, error) {
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

// responderForServiceTasksGet handles the response to the ServiceTasksGet request. The method always
// closes the http.Response Body.
func (c ServiceTaskResourceClient) responderForServiceTasksGet(resp *http.Response) (result ServiceTasksGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

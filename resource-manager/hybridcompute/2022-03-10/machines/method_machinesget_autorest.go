package machines

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MachinesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *Machine
}

type MachinesGetOperationOptions struct {
	Expand *InstanceViewTypes
}

func DefaultMachinesGetOperationOptions() MachinesGetOperationOptions {
	return MachinesGetOperationOptions{}
}

func (o MachinesGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o MachinesGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	return out
}

// MachinesGet ...
func (c MachinesClient) MachinesGet(ctx context.Context, id MachineId, options MachinesGetOperationOptions) (result MachinesGetOperationResponse, err error) {
	req, err := c.preparerForMachinesGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForMachinesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machines.MachinesClient", "MachinesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForMachinesGet prepares the MachinesGet request.
func (c MachinesClient) preparerForMachinesGet(ctx context.Context, id MachineId, options MachinesGetOperationOptions) (*http.Request, error) {
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

// responderForMachinesGet handles the response to the MachinesGet request. The method always
// closes the http.Response Body.
func (c MachinesClient) responderForMachinesGet(resp *http.Response) (result MachinesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

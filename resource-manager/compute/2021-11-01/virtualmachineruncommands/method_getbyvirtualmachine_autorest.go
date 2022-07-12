package virtualmachineruncommands

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetByVirtualMachineOperationResponse struct {
	HttpResponse *http.Response
	Model        *VirtualMachineRunCommand
}

type GetByVirtualMachineOperationOptions struct {
	Expand *string
}

func DefaultGetByVirtualMachineOperationOptions() GetByVirtualMachineOperationOptions {
	return GetByVirtualMachineOperationOptions{}
}

func (o GetByVirtualMachineOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetByVirtualMachineOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	return out
}

// GetByVirtualMachine ...
func (c VirtualMachineRunCommandsClient) GetByVirtualMachine(ctx context.Context, id VirtualMachineRunCommandId, options GetByVirtualMachineOperationOptions) (result GetByVirtualMachineOperationResponse, err error) {
	req, err := c.preparerForGetByVirtualMachine(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineruncommands.VirtualMachineRunCommandsClient", "GetByVirtualMachine", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineruncommands.VirtualMachineRunCommandsClient", "GetByVirtualMachine", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetByVirtualMachine(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineruncommands.VirtualMachineRunCommandsClient", "GetByVirtualMachine", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetByVirtualMachine prepares the GetByVirtualMachine request.
func (c VirtualMachineRunCommandsClient) preparerForGetByVirtualMachine(ctx context.Context, id VirtualMachineRunCommandId, options GetByVirtualMachineOperationOptions) (*http.Request, error) {
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

// responderForGetByVirtualMachine handles the response to the GetByVirtualMachine request. The method always
// closes the http.Response Body.
func (c VirtualMachineRunCommandsClient) responderForGetByVirtualMachine(resp *http.Response) (result GetByVirtualMachineOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

package workloadnetworks

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVirtualMachineOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkloadNetworkVirtualMachine
}

// GetVirtualMachine ...
func (c WorkloadNetworksClient) GetVirtualMachine(ctx context.Context, id DefaultVirtualMachineId) (result GetVirtualMachineOperationResponse, err error) {
	req, err := c.preparerForGetVirtualMachine(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetVirtualMachine", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetVirtualMachine", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetVirtualMachine(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetVirtualMachine", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetVirtualMachine prepares the GetVirtualMachine request.
func (c WorkloadNetworksClient) preparerForGetVirtualMachine(ctx context.Context, id DefaultVirtualMachineId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetVirtualMachine handles the response to the GetVirtualMachine request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForGetVirtualMachine(resp *http.Response) (result GetVirtualMachineOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

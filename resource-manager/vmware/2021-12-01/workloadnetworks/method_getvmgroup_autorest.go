package workloadnetworks

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVMGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkloadNetworkVMGroup
}

// GetVMGroup ...
func (c WorkloadNetworksClient) GetVMGroup(ctx context.Context, id VmGroupId) (result GetVMGroupOperationResponse, err error) {
	req, err := c.preparerForGetVMGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetVMGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetVMGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetVMGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetVMGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetVMGroup prepares the GetVMGroup request.
func (c WorkloadNetworksClient) preparerForGetVMGroup(ctx context.Context, id VmGroupId) (*http.Request, error) {
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

// responderForGetVMGroup handles the response to the GetVMGroup request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForGetVMGroup(resp *http.Response) (result GetVMGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

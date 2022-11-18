package workloadnetworks

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPortMirroringOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkloadNetworkPortMirroring
}

// GetPortMirroring ...
func (c WorkloadNetworksClient) GetPortMirroring(ctx context.Context, id PortMirroringProfileId) (result GetPortMirroringOperationResponse, err error) {
	req, err := c.preparerForGetPortMirroring(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetPortMirroring", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetPortMirroring", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetPortMirroring(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetPortMirroring", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetPortMirroring prepares the GetPortMirroring request.
func (c WorkloadNetworksClient) preparerForGetPortMirroring(ctx context.Context, id PortMirroringProfileId) (*http.Request, error) {
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

// responderForGetPortMirroring handles the response to the GetPortMirroring request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForGetPortMirroring(resp *http.Response) (result GetPortMirroringOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

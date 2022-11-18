package workloadnetworks

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDhcpOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkloadNetworkDhcp
}

// GetDhcp ...
func (c WorkloadNetworksClient) GetDhcp(ctx context.Context, id DhcpConfigurationId) (result GetDhcpOperationResponse, err error) {
	req, err := c.preparerForGetDhcp(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetDhcp", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetDhcp", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDhcp(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetDhcp", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDhcp prepares the GetDhcp request.
func (c WorkloadNetworksClient) preparerForGetDhcp(ctx context.Context, id DhcpConfigurationId) (*http.Request, error) {
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

// responderForGetDhcp handles the response to the GetDhcp request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForGetDhcp(resp *http.Response) (result GetDhcpOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

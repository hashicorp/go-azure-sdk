package workloadnetworks

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPublicIPOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkloadNetworkPublicIP
}

// GetPublicIP ...
func (c WorkloadNetworksClient) GetPublicIP(ctx context.Context, id PublicIPId) (result GetPublicIPOperationResponse, err error) {
	req, err := c.preparerForGetPublicIP(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetPublicIP", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetPublicIP", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetPublicIP(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetPublicIP", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetPublicIP prepares the GetPublicIP request.
func (c WorkloadNetworksClient) preparerForGetPublicIP(ctx context.Context, id PublicIPId) (*http.Request, error) {
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

// responderForGetPublicIP handles the response to the GetPublicIP request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForGetPublicIP(resp *http.Response) (result GetPublicIPOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

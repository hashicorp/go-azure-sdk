package workloadnetworks

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDnsServiceOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkloadNetworkDnsService
}

// GetDnsService ...
func (c WorkloadNetworksClient) GetDnsService(ctx context.Context, id DnsServiceId) (result GetDnsServiceOperationResponse, err error) {
	req, err := c.preparerForGetDnsService(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetDnsService", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetDnsService", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDnsService(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetDnsService", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDnsService prepares the GetDnsService request.
func (c WorkloadNetworksClient) preparerForGetDnsService(ctx context.Context, id DnsServiceId) (*http.Request, error) {
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

// responderForGetDnsService handles the response to the GetDnsService request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForGetDnsService(resp *http.Response) (result GetDnsServiceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

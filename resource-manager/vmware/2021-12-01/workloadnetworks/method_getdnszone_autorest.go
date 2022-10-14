package workloadnetworks

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDnsZoneOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkloadNetworkDnsZone
}

// GetDnsZone ...
func (c WorkloadNetworksClient) GetDnsZone(ctx context.Context, id DnsZoneId) (result GetDnsZoneOperationResponse, err error) {
	req, err := c.preparerForGetDnsZone(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetDnsZone", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetDnsZone", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDnsZone(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetDnsZone", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDnsZone prepares the GetDnsZone request.
func (c WorkloadNetworksClient) preparerForGetDnsZone(ctx context.Context, id DnsZoneId) (*http.Request, error) {
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

// responderForGetDnsZone handles the response to the GetDnsZone request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForGetDnsZone(resp *http.Response) (result GetDnsZoneOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

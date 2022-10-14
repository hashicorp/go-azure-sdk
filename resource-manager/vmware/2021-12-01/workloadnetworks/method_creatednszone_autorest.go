package workloadnetworks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDnsZoneOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateDnsZone ...
func (c WorkloadNetworksClient) CreateDnsZone(ctx context.Context, id DnsZoneId, input WorkloadNetworkDnsZone) (result CreateDnsZoneOperationResponse, err error) {
	req, err := c.preparerForCreateDnsZone(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "CreateDnsZone", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateDnsZone(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "CreateDnsZone", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateDnsZoneThenPoll performs CreateDnsZone then polls until it's completed
func (c WorkloadNetworksClient) CreateDnsZoneThenPoll(ctx context.Context, id DnsZoneId, input WorkloadNetworkDnsZone) error {
	result, err := c.CreateDnsZone(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateDnsZone: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateDnsZone: %+v", err)
	}

	return nil
}

// preparerForCreateDnsZone prepares the CreateDnsZone request.
func (c WorkloadNetworksClient) preparerForCreateDnsZone(ctx context.Context, id DnsZoneId, input WorkloadNetworkDnsZone) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreateDnsZone sends the CreateDnsZone request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForCreateDnsZone(ctx context.Context, req *http.Request) (future CreateDnsZoneOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

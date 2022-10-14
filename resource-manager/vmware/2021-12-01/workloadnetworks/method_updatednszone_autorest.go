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

type UpdateDnsZoneOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// UpdateDnsZone ...
func (c WorkloadNetworksClient) UpdateDnsZone(ctx context.Context, id DnsZoneId, input WorkloadNetworkDnsZone) (result UpdateDnsZoneOperationResponse, err error) {
	req, err := c.preparerForUpdateDnsZone(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "UpdateDnsZone", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUpdateDnsZone(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "UpdateDnsZone", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpdateDnsZoneThenPoll performs UpdateDnsZone then polls until it's completed
func (c WorkloadNetworksClient) UpdateDnsZoneThenPoll(ctx context.Context, id DnsZoneId, input WorkloadNetworkDnsZone) error {
	result, err := c.UpdateDnsZone(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing UpdateDnsZone: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after UpdateDnsZone: %+v", err)
	}

	return nil
}

// preparerForUpdateDnsZone prepares the UpdateDnsZone request.
func (c WorkloadNetworksClient) preparerForUpdateDnsZone(ctx context.Context, id DnsZoneId, input WorkloadNetworkDnsZone) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForUpdateDnsZone sends the UpdateDnsZone request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForUpdateDnsZone(ctx context.Context, req *http.Request) (future UpdateDnsZoneOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

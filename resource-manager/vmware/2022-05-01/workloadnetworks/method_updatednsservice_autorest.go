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

type UpdateDnsServiceOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// UpdateDnsService ...
func (c WorkloadNetworksClient) UpdateDnsService(ctx context.Context, id DnsServiceId, input WorkloadNetworkDnsService) (result UpdateDnsServiceOperationResponse, err error) {
	req, err := c.preparerForUpdateDnsService(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "UpdateDnsService", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUpdateDnsService(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "UpdateDnsService", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpdateDnsServiceThenPoll performs UpdateDnsService then polls until it's completed
func (c WorkloadNetworksClient) UpdateDnsServiceThenPoll(ctx context.Context, id DnsServiceId, input WorkloadNetworkDnsService) error {
	result, err := c.UpdateDnsService(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing UpdateDnsService: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after UpdateDnsService: %+v", err)
	}

	return nil
}

// preparerForUpdateDnsService prepares the UpdateDnsService request.
func (c WorkloadNetworksClient) preparerForUpdateDnsService(ctx context.Context, id DnsServiceId, input WorkloadNetworkDnsService) (*http.Request, error) {
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

// senderForUpdateDnsService sends the UpdateDnsService request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForUpdateDnsService(ctx context.Context, req *http.Request) (future UpdateDnsServiceOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

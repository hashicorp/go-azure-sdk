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

type DeleteDnsServiceOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeleteDnsService ...
func (c WorkloadNetworksClient) DeleteDnsService(ctx context.Context, id DnsServiceId) (result DeleteDnsServiceOperationResponse, err error) {
	req, err := c.preparerForDeleteDnsService(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "DeleteDnsService", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteDnsService(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "DeleteDnsService", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteDnsServiceThenPoll performs DeleteDnsService then polls until it's completed
func (c WorkloadNetworksClient) DeleteDnsServiceThenPoll(ctx context.Context, id DnsServiceId) error {
	result, err := c.DeleteDnsService(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteDnsService: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteDnsService: %+v", err)
	}

	return nil
}

// preparerForDeleteDnsService prepares the DeleteDnsService request.
func (c WorkloadNetworksClient) preparerForDeleteDnsService(ctx context.Context, id DnsServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeleteDnsService sends the DeleteDnsService request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForDeleteDnsService(ctx context.Context, req *http.Request) (future DeleteDnsServiceOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

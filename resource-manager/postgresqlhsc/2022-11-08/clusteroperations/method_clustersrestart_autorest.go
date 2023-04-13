package clusteroperations

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

type ClustersRestartOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ClustersRestart ...
func (c ClusterOperationsClient) ClustersRestart(ctx context.Context, id ServerGroupsv2Id) (result ClustersRestartOperationResponse, err error) {
	req, err := c.preparerForClustersRestart(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "clusteroperations.ClusterOperationsClient", "ClustersRestart", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForClustersRestart(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "clusteroperations.ClusterOperationsClient", "ClustersRestart", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ClustersRestartThenPoll performs ClustersRestart then polls until it's completed
func (c ClusterOperationsClient) ClustersRestartThenPoll(ctx context.Context, id ServerGroupsv2Id) error {
	result, err := c.ClustersRestart(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ClustersRestart: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ClustersRestart: %+v", err)
	}

	return nil
}

// preparerForClustersRestart prepares the ClustersRestart request.
func (c ClusterOperationsClient) preparerForClustersRestart(ctx context.Context, id ServerGroupsv2Id) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/restart", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForClustersRestart sends the ClustersRestart request. The method will close the
// http.Response Body if it receives an error.
func (c ClusterOperationsClient) senderForClustersRestart(ctx context.Context, req *http.Request) (future ClustersRestartOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

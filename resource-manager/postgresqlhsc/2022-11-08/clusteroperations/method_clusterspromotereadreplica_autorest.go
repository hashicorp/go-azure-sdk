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

type ClustersPromoteReadReplicaOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ClustersPromoteReadReplica ...
func (c ClusterOperationsClient) ClustersPromoteReadReplica(ctx context.Context, id ServerGroupsv2Id) (result ClustersPromoteReadReplicaOperationResponse, err error) {
	req, err := c.preparerForClustersPromoteReadReplica(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "clusteroperations.ClusterOperationsClient", "ClustersPromoteReadReplica", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForClustersPromoteReadReplica(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "clusteroperations.ClusterOperationsClient", "ClustersPromoteReadReplica", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ClustersPromoteReadReplicaThenPoll performs ClustersPromoteReadReplica then polls until it's completed
func (c ClusterOperationsClient) ClustersPromoteReadReplicaThenPoll(ctx context.Context, id ServerGroupsv2Id) error {
	result, err := c.ClustersPromoteReadReplica(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ClustersPromoteReadReplica: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ClustersPromoteReadReplica: %+v", err)
	}

	return nil
}

// preparerForClustersPromoteReadReplica prepares the ClustersPromoteReadReplica request.
func (c ClusterOperationsClient) preparerForClustersPromoteReadReplica(ctx context.Context, id ServerGroupsv2Id) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/promote", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForClustersPromoteReadReplica sends the ClustersPromoteReadReplica request. The method will close the
// http.Response Body if it receives an error.
func (c ClusterOperationsClient) senderForClustersPromoteReadReplica(ctx context.Context, req *http.Request) (future ClustersPromoteReadReplicaOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

package clusteroperations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClustersPromoteReadReplicaOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ClustersPromoteReadReplica ...
func (c ClusterOperationsClient) ClustersPromoteReadReplica(ctx context.Context, id ServerGroupsv2Id) (result ClustersPromoteReadReplicaOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
		},
		HttpMethod: http.MethodPost,

		Path: fmt.Sprintf("%s/promote", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
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

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ClustersPromoteReadReplica: %+v", err)
	}

	return nil
}

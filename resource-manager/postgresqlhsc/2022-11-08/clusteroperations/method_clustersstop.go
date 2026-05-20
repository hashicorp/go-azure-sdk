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

type ClustersStopOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ClustersStop ...
func (c ClusterOperationsClient) ClustersStop(ctx context.Context, id ServerGroupsv2Id) (result ClustersStopOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/stop", id.ID()),
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

// ClustersStopThenPoll performs ClustersStop then polls until it's completed
func (c ClusterOperationsClient) ClustersStopThenPoll(ctx context.Context, id ServerGroupsv2Id) error {
	return c.ClustersStopCallbackThenPoll(ctx, id, nil)
}

// ClustersStopCallbackThenPoll performs ClustersStop, runs the optional callback function, then polls until it's completed
func (c ClusterOperationsClient) ClustersStopCallbackThenPoll(ctx context.Context, id ServerGroupsv2Id, callback func() error) error {
	result, err := c.ClustersStop(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ClustersStop: %+v", err)
	}

	if callback != nil {
		if err := callback(); err != nil {
			return fmt.Errorf("executing callback function: %+v", err)
		}
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ClustersStop: %+v", err)
	}

	return nil
}

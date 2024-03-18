package vmwares

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

type WorkloadNetworksCreatePortMirroringOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *WorkloadNetworkPortMirroring
}

// WorkloadNetworksCreatePortMirroring ...
func (c VMwaresClient) WorkloadNetworksCreatePortMirroring(ctx context.Context, id PortMirroringProfileId, input WorkloadNetworkPortMirroring) (result WorkloadNetworksCreatePortMirroringOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

// WorkloadNetworksCreatePortMirroringThenPoll performs WorkloadNetworksCreatePortMirroring then polls until it's completed
func (c VMwaresClient) WorkloadNetworksCreatePortMirroringThenPoll(ctx context.Context, id PortMirroringProfileId, input WorkloadNetworkPortMirroring) error {
	result, err := c.WorkloadNetworksCreatePortMirroring(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WorkloadNetworksCreatePortMirroring: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after WorkloadNetworksCreatePortMirroring: %+v", err)
	}

	return nil
}

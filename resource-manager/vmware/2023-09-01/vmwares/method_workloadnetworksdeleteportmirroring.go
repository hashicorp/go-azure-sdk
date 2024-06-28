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

type WorkloadNetworksDeletePortMirroringOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// WorkloadNetworksDeletePortMirroring ...
func (c VMwaresClient) WorkloadNetworksDeletePortMirroring(ctx context.Context, id PortMirroringProfileId) (result WorkloadNetworksDeletePortMirroringOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod: http.MethodDelete,

		Path: id.ID(),
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

// WorkloadNetworksDeletePortMirroringThenPoll performs WorkloadNetworksDeletePortMirroring then polls until it's completed
func (c VMwaresClient) WorkloadNetworksDeletePortMirroringThenPoll(ctx context.Context, id PortMirroringProfileId) error {
	result, err := c.WorkloadNetworksDeletePortMirroring(ctx, id)
	if err != nil {
		return fmt.Errorf("performing WorkloadNetworksDeletePortMirroring: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after WorkloadNetworksDeletePortMirroring: %+v", err)
	}

	return nil
}

package databoxedgedevices

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

type DevicesDownloadUpdatesOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// DevicesDownloadUpdates ...
func (c DataBoxEdgeDevicesClient) DevicesDownloadUpdates(ctx context.Context, id DataBoxEdgeDeviceId) (result DevicesDownloadUpdatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/downloadUpdates", id.ID()),
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

// DevicesDownloadUpdatesThenPoll performs DevicesDownloadUpdates then polls until it's completed
func (c DataBoxEdgeDevicesClient) DevicesDownloadUpdatesThenPoll(ctx context.Context, id DataBoxEdgeDeviceId) error {
	result, err := c.DevicesDownloadUpdates(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DevicesDownloadUpdates: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DevicesDownloadUpdates: %+v", err)
	}

	return nil
}

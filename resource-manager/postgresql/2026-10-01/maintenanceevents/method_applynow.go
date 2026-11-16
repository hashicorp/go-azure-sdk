package maintenanceevents

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

type ApplyNowOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *MaintenanceEventActionResponse
}

// ApplyNow ...
func (c MaintenanceEventsClient) ApplyNow(ctx context.Context, id MaintenanceEventId) (result ApplyNowOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/applyNow", id.ID()),
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

// ApplyNowThenPoll performs ApplyNow then polls until it's completed
func (c MaintenanceEventsClient) ApplyNowThenPoll(ctx context.Context, id MaintenanceEventId) error {
	return c.ApplyNowCallbackThenPoll(ctx, id, nil)
}

// ApplyNowCallbackThenPoll performs ApplyNow, runs the optional callback function, then polls until it's completed
func (c MaintenanceEventsClient) ApplyNowCallbackThenPoll(ctx context.Context, id MaintenanceEventId, callback func() error) error {
	result, err := c.ApplyNow(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ApplyNow: %+v", err)
	}

	if callback != nil {
		if err := callback(); err != nil {
			return fmt.Errorf("executing callback function: %+v", err)
		}
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ApplyNow: %+v", err)
	}

	return nil
}

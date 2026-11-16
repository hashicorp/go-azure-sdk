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

type RescheduleOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *MaintenanceEventActionResponse
}

// Reschedule ...
func (c MaintenanceEventsClient) Reschedule(ctx context.Context, id MaintenanceEventId, input MaintenanceEventRescheduleRequest) (result RescheduleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/reschedule", id.ID()),
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

// RescheduleThenPoll performs Reschedule then polls until it's completed
func (c MaintenanceEventsClient) RescheduleThenPoll(ctx context.Context, id MaintenanceEventId, input MaintenanceEventRescheduleRequest) error {
	return c.RescheduleCallbackThenPoll(ctx, id, input, nil)
}

// RescheduleCallbackThenPoll performs Reschedule, runs the optional callback function, then polls until it's completed
func (c MaintenanceEventsClient) RescheduleCallbackThenPoll(ctx context.Context, id MaintenanceEventId, input MaintenanceEventRescheduleRequest, callback func() error) error {
	result, err := c.Reschedule(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Reschedule: %+v", err)
	}

	if callback != nil {
		if err := callback(); err != nil {
			return fmt.Errorf("executing callback function: %+v", err)
		}
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after Reschedule: %+v", err)
	}

	return nil
}

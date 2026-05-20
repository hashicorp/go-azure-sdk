package storagetasks

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

type StopAllAssignmentsOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// StopAllAssignments ...
func (c StorageTasksClient) StopAllAssignments(ctx context.Context, id StorageTaskId) (result StopAllAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/stopAllAssignments", id.ID()),
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

// StopAllAssignmentsThenPoll performs StopAllAssignments then polls until it's completed
func (c StorageTasksClient) StopAllAssignmentsThenPoll(ctx context.Context, id StorageTaskId) error {
	return c.StopAllAssignmentsCallbackThenPoll(ctx, id, nil)
}

// StopAllAssignmentsCallbackThenPoll performs StopAllAssignments, runs the optional callback function, then polls until it's completed
func (c StorageTasksClient) StopAllAssignmentsCallbackThenPoll(ctx context.Context, id StorageTaskId, callback func() error) error {
	result, err := c.StopAllAssignments(ctx, id)
	if err != nil {
		return fmt.Errorf("performing StopAllAssignments: %+v", err)
	}

	if callback != nil {
		if err := callback(); err != nil {
			return fmt.Errorf("executing callback function: %+v", err)
		}
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after StopAllAssignments: %+v", err)
	}

	return nil
}

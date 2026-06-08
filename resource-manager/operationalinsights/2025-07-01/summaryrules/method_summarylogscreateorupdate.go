package summaryrules

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

type SummaryLogsCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SummaryLogs
}

// SummaryLogsCreateOrUpdate ...
func (c SummaryRulesClient) SummaryLogsCreateOrUpdate(ctx context.Context, id SummaryLogId, input SummaryLogs) (result SummaryLogsCreateOrUpdateOperationResponse, err error) {
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

// SummaryLogsCreateOrUpdateThenPoll performs SummaryLogsCreateOrUpdate then polls until it's completed
func (c SummaryRulesClient) SummaryLogsCreateOrUpdateThenPoll(ctx context.Context, id SummaryLogId, input SummaryLogs) error {
	return c.SummaryLogsCreateOrUpdateCallbackThenPoll(ctx, id, input, nil)
}

// SummaryLogsCreateOrUpdateCallbackThenPoll performs SummaryLogsCreateOrUpdate, runs the optional callback function, then polls until it's completed
func (c SummaryRulesClient) SummaryLogsCreateOrUpdateCallbackThenPoll(ctx context.Context, id SummaryLogId, input SummaryLogs, callback func() error) error {
	result, err := c.SummaryLogsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SummaryLogsCreateOrUpdate: %+v", err)
	}

	if callback != nil {
		if err := callback(); err != nil {
			return fmt.Errorf("executing callback function: %+v", err)
		}
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after SummaryLogsCreateOrUpdate: %+v", err)
	}

	return nil
}

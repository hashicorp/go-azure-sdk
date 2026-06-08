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

type SummaryLogsRetryBinOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// SummaryLogsRetryBin ...
func (c SummaryRulesClient) SummaryLogsRetryBin(ctx context.Context, id SummaryLogId, input SummaryLogsRetryBin) (result SummaryLogsRetryBinOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/retrybin", id.ID()),
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

// SummaryLogsRetryBinThenPoll performs SummaryLogsRetryBin then polls until it's completed
func (c SummaryRulesClient) SummaryLogsRetryBinThenPoll(ctx context.Context, id SummaryLogId, input SummaryLogsRetryBin) error {
	return c.SummaryLogsRetryBinCallbackThenPoll(ctx, id, input, nil)
}

// SummaryLogsRetryBinCallbackThenPoll performs SummaryLogsRetryBin, runs the optional callback function, then polls until it's completed
func (c SummaryRulesClient) SummaryLogsRetryBinCallbackThenPoll(ctx context.Context, id SummaryLogId, input SummaryLogsRetryBin, callback func() error) error {
	result, err := c.SummaryLogsRetryBin(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SummaryLogsRetryBin: %+v", err)
	}

	if callback != nil {
		if err := callback(); err != nil {
			return fmt.Errorf("executing callback function: %+v", err)
		}
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after SummaryLogsRetryBin: %+v", err)
	}

	return nil
}

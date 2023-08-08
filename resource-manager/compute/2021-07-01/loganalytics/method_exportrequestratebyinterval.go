package loganalytics

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

type ExportRequestRateByIntervalOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ExportRequestRateByInterval ...
func (c LogAnalyticsClient) ExportRequestRateByInterval(ctx context.Context, id LocationId, input RequestRateByIntervalInput) (result ExportRequestRateByIntervalOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/logAnalytics/apiAccess/getRequestRateByInterval", id.ID()),
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

// ExportRequestRateByIntervalThenPoll performs ExportRequestRateByInterval then polls until it's completed
func (c LogAnalyticsClient) ExportRequestRateByIntervalThenPoll(ctx context.Context, id LocationId, input RequestRateByIntervalInput) error {
	result, err := c.ExportRequestRateByInterval(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ExportRequestRateByInterval: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ExportRequestRateByInterval: %+v", err)
	}

	return nil
}

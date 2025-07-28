package computerps

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

type LogAnalyticsExportThrottledRequestsOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *LogAnalyticsOperationResult
}

// LogAnalyticsExportThrottledRequests ...
func (c ComputeRPSClient) LogAnalyticsExportThrottledRequests(ctx context.Context, id LocationId, input LogAnalyticsInputBase) (result LogAnalyticsExportThrottledRequestsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/logAnalytics/apiAccess/getThrottledRequests", id.ID()),
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

// LogAnalyticsExportThrottledRequestsThenPoll performs LogAnalyticsExportThrottledRequests then polls until it's completed
func (c ComputeRPSClient) LogAnalyticsExportThrottledRequestsThenPoll(ctx context.Context, id LocationId, input LogAnalyticsInputBase) error {
	result, err := c.LogAnalyticsExportThrottledRequests(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing LogAnalyticsExportThrottledRequests: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after LogAnalyticsExportThrottledRequests: %+v", err)
	}

	return nil
}

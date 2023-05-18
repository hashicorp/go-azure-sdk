package costdetails

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GenerateCostDetailsReportCreateOperationOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// GenerateCostDetailsReportCreateOperation ...
func (c CostDetailsClient) GenerateCostDetailsReportCreateOperation(ctx context.Context, id commonids.ScopeId, input GenerateCostDetailsReportRequestDefinition) (result GenerateCostDetailsReportCreateOperationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/providers/Microsoft.CostManagement/generateCostDetailsReport", id.ID()),
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

// GenerateCostDetailsReportCreateOperationThenPoll performs GenerateCostDetailsReportCreateOperation then polls until it's completed
func (c CostDetailsClient) GenerateCostDetailsReportCreateOperationThenPoll(ctx context.Context, id commonids.ScopeId, input GenerateCostDetailsReportRequestDefinition) error {
	result, err := c.GenerateCostDetailsReportCreateOperation(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing GenerateCostDetailsReportCreateOperation: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after GenerateCostDetailsReportCreateOperation: %+v", err)
	}

	return nil
}

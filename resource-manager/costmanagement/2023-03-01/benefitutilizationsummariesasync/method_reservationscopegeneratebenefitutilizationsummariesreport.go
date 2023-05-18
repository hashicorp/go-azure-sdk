package benefitutilizationsummariesasync

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

type ReservationScopeGenerateBenefitUtilizationSummariesReportOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ReservationScopeGenerateBenefitUtilizationSummariesReport ...
func (c BenefitUtilizationSummariesAsyncClient) ReservationScopeGenerateBenefitUtilizationSummariesReport(ctx context.Context, id ReservationId, input BenefitUtilizationSummariesRequest) (result ReservationScopeGenerateBenefitUtilizationSummariesReportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/providers/Microsoft.CostManagement/generateBenefitUtilizationSummariesReport", id.ID()),
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

// ReservationScopeGenerateBenefitUtilizationSummariesReportThenPoll performs ReservationScopeGenerateBenefitUtilizationSummariesReport then polls until it's completed
func (c BenefitUtilizationSummariesAsyncClient) ReservationScopeGenerateBenefitUtilizationSummariesReportThenPoll(ctx context.Context, id ReservationId, input BenefitUtilizationSummariesRequest) error {
	result, err := c.ReservationScopeGenerateBenefitUtilizationSummariesReport(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ReservationScopeGenerateBenefitUtilizationSummariesReport: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ReservationScopeGenerateBenefitUtilizationSummariesReport: %+v", err)
	}

	return nil
}

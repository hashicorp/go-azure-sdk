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

type GenerateBenefitUtilizationSummariesReportGenerateByBillingProfileOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *BenefitUtilizationSummariesOperationStatus
}

// GenerateBenefitUtilizationSummariesReportGenerateByBillingProfile ...
func (c BenefitUtilizationSummariesAsyncClient) GenerateBenefitUtilizationSummariesReportGenerateByBillingProfile(ctx context.Context, id BillingProfileId, input BenefitUtilizationSummariesRequest) (result GenerateBenefitUtilizationSummariesReportGenerateByBillingProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
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

// GenerateBenefitUtilizationSummariesReportGenerateByBillingProfileThenPoll performs GenerateBenefitUtilizationSummariesReportGenerateByBillingProfile then polls until it's completed
func (c BenefitUtilizationSummariesAsyncClient) GenerateBenefitUtilizationSummariesReportGenerateByBillingProfileThenPoll(ctx context.Context, id BillingProfileId, input BenefitUtilizationSummariesRequest) error {
	return c.GenerateBenefitUtilizationSummariesReportGenerateByBillingProfileCallbackThenPoll(ctx, id, input, nil)
}

// GenerateBenefitUtilizationSummariesReportGenerateByBillingProfileCallbackThenPoll performs GenerateBenefitUtilizationSummariesReportGenerateByBillingProfile, runs the optional callback function, then polls until it's completed
func (c BenefitUtilizationSummariesAsyncClient) GenerateBenefitUtilizationSummariesReportGenerateByBillingProfileCallbackThenPoll(ctx context.Context, id BillingProfileId, input BenefitUtilizationSummariesRequest, callback func() error) error {
	result, err := c.GenerateBenefitUtilizationSummariesReportGenerateByBillingProfile(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing GenerateBenefitUtilizationSummariesReportGenerateByBillingProfile: %+v", err)
	}

	if callback != nil {
		if err := callback(); err != nil {
			return fmt.Errorf("executing callback function: %+v", err)
		}
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after GenerateBenefitUtilizationSummariesReportGenerateByBillingProfile: %+v", err)
	}

	return nil
}

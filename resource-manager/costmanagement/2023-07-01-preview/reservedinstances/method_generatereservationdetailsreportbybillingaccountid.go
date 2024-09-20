package reservedinstances

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

type GenerateReservationDetailsReportByBillingAccountIdOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *OperationStatus
}

type GenerateReservationDetailsReportByBillingAccountIdOperationOptions struct {
	EndDate   *string
	StartDate *string
}

func DefaultGenerateReservationDetailsReportByBillingAccountIdOperationOptions() GenerateReservationDetailsReportByBillingAccountIdOperationOptions {
	return GenerateReservationDetailsReportByBillingAccountIdOperationOptions{}
}

func (o GenerateReservationDetailsReportByBillingAccountIdOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GenerateReservationDetailsReportByBillingAccountIdOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GenerateReservationDetailsReportByBillingAccountIdOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndDate != nil {
		out.Append("endDate", fmt.Sprintf("%v", *o.EndDate))
	}
	if o.StartDate != nil {
		out.Append("startDate", fmt.Sprintf("%v", *o.StartDate))
	}
	return &out
}

// GenerateReservationDetailsReportByBillingAccountId ...
func (c ReservedInstancesClient) GenerateReservationDetailsReportByBillingAccountId(ctx context.Context, id BillingAccountId, options GenerateReservationDetailsReportByBillingAccountIdOperationOptions) (result GenerateReservationDetailsReportByBillingAccountIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/providers/Microsoft.CostManagement/generateReservationDetailsReport", id.ID()),
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

// GenerateReservationDetailsReportByBillingAccountIdThenPoll performs GenerateReservationDetailsReportByBillingAccountId then polls until it's completed
func (c ReservedInstancesClient) GenerateReservationDetailsReportByBillingAccountIdThenPoll(ctx context.Context, id BillingAccountId, options GenerateReservationDetailsReportByBillingAccountIdOperationOptions) error {
	result, err := c.GenerateReservationDetailsReportByBillingAccountId(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing GenerateReservationDetailsReportByBillingAccountId: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after GenerateReservationDetailsReportByBillingAccountId: %+v", err)
	}

	return nil
}

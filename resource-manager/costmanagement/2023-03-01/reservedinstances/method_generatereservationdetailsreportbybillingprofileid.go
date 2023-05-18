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

type GenerateReservationDetailsReportByBillingProfileIdOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type GenerateReservationDetailsReportByBillingProfileIdOperationOptions struct {
	EndDate   *string
	StartDate *string
}

func DefaultGenerateReservationDetailsReportByBillingProfileIdOperationOptions() GenerateReservationDetailsReportByBillingProfileIdOperationOptions {
	return GenerateReservationDetailsReportByBillingProfileIdOperationOptions{}
}

func (o GenerateReservationDetailsReportByBillingProfileIdOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GenerateReservationDetailsReportByBillingProfileIdOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GenerateReservationDetailsReportByBillingProfileIdOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndDate != nil {
		out.Append("endDate", fmt.Sprintf("%v", *o.EndDate))
	}
	if o.StartDate != nil {
		out.Append("startDate", fmt.Sprintf("%v", *o.StartDate))
	}
	return &out
}

// GenerateReservationDetailsReportByBillingProfileId ...
func (c ReservedInstancesClient) GenerateReservationDetailsReportByBillingProfileId(ctx context.Context, id BillingProfileId, options GenerateReservationDetailsReportByBillingProfileIdOperationOptions) (result GenerateReservationDetailsReportByBillingProfileIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/providers/Microsoft.CostManagement/generateReservationDetailsReport", id.ID()),
		OptionsObject: options,
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

// GenerateReservationDetailsReportByBillingProfileIdThenPoll performs GenerateReservationDetailsReportByBillingProfileId then polls until it's completed
func (c ReservedInstancesClient) GenerateReservationDetailsReportByBillingProfileIdThenPoll(ctx context.Context, id BillingProfileId, options GenerateReservationDetailsReportByBillingProfileIdOperationOptions) error {
	result, err := c.GenerateReservationDetailsReportByBillingProfileId(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing GenerateReservationDetailsReportByBillingProfileId: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after GenerateReservationDetailsReportByBillingProfileId: %+v", err)
	}

	return nil
}

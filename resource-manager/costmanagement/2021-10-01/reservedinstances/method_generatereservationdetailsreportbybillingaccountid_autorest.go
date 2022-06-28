package reservedinstances

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GenerateReservationDetailsReportByBillingAccountIdOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type GenerateReservationDetailsReportByBillingAccountIdOperationOptions struct {
	EndDate   *string
	StartDate *string
}

func DefaultGenerateReservationDetailsReportByBillingAccountIdOperationOptions() GenerateReservationDetailsReportByBillingAccountIdOperationOptions {
	return GenerateReservationDetailsReportByBillingAccountIdOperationOptions{}
}

func (o GenerateReservationDetailsReportByBillingAccountIdOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GenerateReservationDetailsReportByBillingAccountIdOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.EndDate != nil {
		out["endDate"] = *o.EndDate
	}

	if o.StartDate != nil {
		out["startDate"] = *o.StartDate
	}

	return out
}

// GenerateReservationDetailsReportByBillingAccountId ...
func (c ReservedInstancesClient) GenerateReservationDetailsReportByBillingAccountId(ctx context.Context, id BillingAccountId, options GenerateReservationDetailsReportByBillingAccountIdOperationOptions) (result GenerateReservationDetailsReportByBillingAccountIdOperationResponse, err error) {
	req, err := c.preparerForGenerateReservationDetailsReportByBillingAccountId(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservedinstances.ReservedInstancesClient", "GenerateReservationDetailsReportByBillingAccountId", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForGenerateReservationDetailsReportByBillingAccountId(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservedinstances.ReservedInstancesClient", "GenerateReservationDetailsReportByBillingAccountId", result.HttpResponse, "Failure sending request")
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

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after GenerateReservationDetailsReportByBillingAccountId: %+v", err)
	}

	return nil
}

// preparerForGenerateReservationDetailsReportByBillingAccountId prepares the GenerateReservationDetailsReportByBillingAccountId request.
func (c ReservedInstancesClient) preparerForGenerateReservationDetailsReportByBillingAccountId(ctx context.Context, id BillingAccountId, options GenerateReservationDetailsReportByBillingAccountIdOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.CostManagement/generateReservationDetailsReport", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForGenerateReservationDetailsReportByBillingAccountId sends the GenerateReservationDetailsReportByBillingAccountId request. The method will close the
// http.Response Body if it receives an error.
func (c ReservedInstancesClient) senderForGenerateReservationDetailsReportByBillingAccountId(ctx context.Context, req *http.Request) (future GenerateReservationDetailsReportByBillingAccountIdOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}

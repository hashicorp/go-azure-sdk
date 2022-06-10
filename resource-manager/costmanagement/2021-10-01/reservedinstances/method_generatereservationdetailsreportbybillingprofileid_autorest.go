package reservedinstances

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

type GenerateReservationDetailsReportByBillingProfileIdOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type GenerateReservationDetailsReportByBillingProfileIdOperationOptions struct {
	EndDate   *string
	StartDate *string
}

func DefaultGenerateReservationDetailsReportByBillingProfileIdOperationOptions() GenerateReservationDetailsReportByBillingProfileIdOperationOptions {
	return GenerateReservationDetailsReportByBillingProfileIdOperationOptions{}
}

func (o GenerateReservationDetailsReportByBillingProfileIdOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GenerateReservationDetailsReportByBillingProfileIdOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.EndDate != nil {
		out["endDate"] = *o.EndDate
	}

	if o.StartDate != nil {
		out["startDate"] = *o.StartDate
	}

	return out
}

// GenerateReservationDetailsReportByBillingProfileId ...
func (c ReservedInstancesClient) GenerateReservationDetailsReportByBillingProfileId(ctx context.Context, id BillingProfileId, options GenerateReservationDetailsReportByBillingProfileIdOperationOptions) (result GenerateReservationDetailsReportByBillingProfileIdOperationResponse, err error) {
	req, err := c.preparerForGenerateReservationDetailsReportByBillingProfileId(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservedinstances.ReservedInstancesClient", "GenerateReservationDetailsReportByBillingProfileId", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForGenerateReservationDetailsReportByBillingProfileId(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "reservedinstances.ReservedInstancesClient", "GenerateReservationDetailsReportByBillingProfileId", result.HttpResponse, "Failure sending request")
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

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after GenerateReservationDetailsReportByBillingProfileId: %+v", err)
	}

	return nil
}

// preparerForGenerateReservationDetailsReportByBillingProfileId prepares the GenerateReservationDetailsReportByBillingProfileId request.
func (c ReservedInstancesClient) preparerForGenerateReservationDetailsReportByBillingProfileId(ctx context.Context, id BillingProfileId, options GenerateReservationDetailsReportByBillingProfileIdOperationOptions) (*http.Request, error) {
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

// senderForGenerateReservationDetailsReportByBillingProfileId sends the GenerateReservationDetailsReportByBillingProfileId request. The method will close the
// http.Response Body if it receives an error.
func (c ReservedInstancesClient) senderForGenerateReservationDetailsReportByBillingProfileId(ctx context.Context, req *http.Request) (future GenerateReservationDetailsReportByBillingProfileIdOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}
	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}

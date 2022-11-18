package costdetails

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GenerateCostDetailsReportCreateOperationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// GenerateCostDetailsReportCreateOperation ...
func (c CostDetailsClient) GenerateCostDetailsReportCreateOperation(ctx context.Context, id commonids.ScopeId, input GenerateCostDetailsReportRequestDefinition) (result GenerateCostDetailsReportCreateOperationOperationResponse, err error) {
	req, err := c.preparerForGenerateCostDetailsReportCreateOperation(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costdetails.CostDetailsClient", "GenerateCostDetailsReportCreateOperation", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForGenerateCostDetailsReportCreateOperation(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "costdetails.CostDetailsClient", "GenerateCostDetailsReportCreateOperation", result.HttpResponse, "Failure sending request")
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

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after GenerateCostDetailsReportCreateOperation: %+v", err)
	}

	return nil
}

// preparerForGenerateCostDetailsReportCreateOperation prepares the GenerateCostDetailsReportCreateOperation request.
func (c CostDetailsClient) preparerForGenerateCostDetailsReportCreateOperation(ctx context.Context, id commonids.ScopeId, input GenerateCostDetailsReportRequestDefinition) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.CostManagement/generateCostDetailsReport", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForGenerateCostDetailsReportCreateOperation sends the GenerateCostDetailsReportCreateOperation request. The method will close the
// http.Response Body if it receives an error.
func (c CostDetailsClient) senderForGenerateCostDetailsReportCreateOperation(ctx context.Context, req *http.Request) (future GenerateCostDetailsReportCreateOperationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

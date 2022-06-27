package usagedetails

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

type GenerateDetailedCostReportCreateOperationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// GenerateDetailedCostReportCreateOperation ...
func (c UsageDetailsClient) GenerateDetailedCostReportCreateOperation(ctx context.Context, id commonids.ScopeId, input GenerateDetailedCostReportDefinition) (result GenerateDetailedCostReportCreateOperationOperationResponse, err error) {
	req, err := c.preparerForGenerateDetailedCostReportCreateOperation(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "usagedetails.UsageDetailsClient", "GenerateDetailedCostReportCreateOperation", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForGenerateDetailedCostReportCreateOperation(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "usagedetails.UsageDetailsClient", "GenerateDetailedCostReportCreateOperation", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// GenerateDetailedCostReportCreateOperationThenPoll performs GenerateDetailedCostReportCreateOperation then polls until it's completed
func (c UsageDetailsClient) GenerateDetailedCostReportCreateOperationThenPoll(ctx context.Context, id commonids.ScopeId, input GenerateDetailedCostReportDefinition) error {
	result, err := c.GenerateDetailedCostReportCreateOperation(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing GenerateDetailedCostReportCreateOperation: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after GenerateDetailedCostReportCreateOperation: %+v", err)
	}

	return nil
}

// preparerForGenerateDetailedCostReportCreateOperation prepares the GenerateDetailedCostReportCreateOperation request.
func (c UsageDetailsClient) preparerForGenerateDetailedCostReportCreateOperation(ctx context.Context, id commonids.ScopeId, input GenerateDetailedCostReportDefinition) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.CostManagement/generateDetailedCostReport", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForGenerateDetailedCostReportCreateOperation sends the GenerateDetailedCostReportCreateOperation request. The method will close the
// http.Response Body if it receives an error.
func (c UsageDetailsClient) senderForGenerateDetailedCostReportCreateOperation(ctx context.Context, req *http.Request) (future GenerateDetailedCostReportCreateOperationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)

	return
}

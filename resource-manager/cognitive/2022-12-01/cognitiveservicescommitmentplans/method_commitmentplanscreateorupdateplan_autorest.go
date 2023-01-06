package cognitiveservicescommitmentplans

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

type CommitmentPlansCreateOrUpdatePlanOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CommitmentPlansCreateOrUpdatePlan ...
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansCreateOrUpdatePlan(ctx context.Context, id CommitmentPlanId, input CommitmentPlan) (result CommitmentPlansCreateOrUpdatePlanOperationResponse, err error) {
	req, err := c.preparerForCommitmentPlansCreateOrUpdatePlan(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansCreateOrUpdatePlan", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCommitmentPlansCreateOrUpdatePlan(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansCreateOrUpdatePlan", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CommitmentPlansCreateOrUpdatePlanThenPoll performs CommitmentPlansCreateOrUpdatePlan then polls until it's completed
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansCreateOrUpdatePlanThenPoll(ctx context.Context, id CommitmentPlanId, input CommitmentPlan) error {
	result, err := c.CommitmentPlansCreateOrUpdatePlan(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CommitmentPlansCreateOrUpdatePlan: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CommitmentPlansCreateOrUpdatePlan: %+v", err)
	}

	return nil
}

// preparerForCommitmentPlansCreateOrUpdatePlan prepares the CommitmentPlansCreateOrUpdatePlan request.
func (c CognitiveServicesCommitmentPlansClient) preparerForCommitmentPlansCreateOrUpdatePlan(ctx context.Context, id CommitmentPlanId, input CommitmentPlan) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCommitmentPlansCreateOrUpdatePlan sends the CommitmentPlansCreateOrUpdatePlan request. The method will close the
// http.Response Body if it receives an error.
func (c CognitiveServicesCommitmentPlansClient) senderForCommitmentPlansCreateOrUpdatePlan(ctx context.Context, req *http.Request) (future CommitmentPlansCreateOrUpdatePlanOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

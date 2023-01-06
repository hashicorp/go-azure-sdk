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

type CommitmentPlansDeletePlanOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CommitmentPlansDeletePlan ...
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansDeletePlan(ctx context.Context, id CommitmentPlanId) (result CommitmentPlansDeletePlanOperationResponse, err error) {
	req, err := c.preparerForCommitmentPlansDeletePlan(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansDeletePlan", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCommitmentPlansDeletePlan(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansDeletePlan", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CommitmentPlansDeletePlanThenPoll performs CommitmentPlansDeletePlan then polls until it's completed
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansDeletePlanThenPoll(ctx context.Context, id CommitmentPlanId) error {
	result, err := c.CommitmentPlansDeletePlan(ctx, id)
	if err != nil {
		return fmt.Errorf("performing CommitmentPlansDeletePlan: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CommitmentPlansDeletePlan: %+v", err)
	}

	return nil
}

// preparerForCommitmentPlansDeletePlan prepares the CommitmentPlansDeletePlan request.
func (c CognitiveServicesCommitmentPlansClient) preparerForCommitmentPlansDeletePlan(ctx context.Context, id CommitmentPlanId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCommitmentPlansDeletePlan sends the CommitmentPlansDeletePlan request. The method will close the
// http.Response Body if it receives an error.
func (c CognitiveServicesCommitmentPlansClient) senderForCommitmentPlansDeletePlan(ctx context.Context, req *http.Request) (future CommitmentPlansDeletePlanOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

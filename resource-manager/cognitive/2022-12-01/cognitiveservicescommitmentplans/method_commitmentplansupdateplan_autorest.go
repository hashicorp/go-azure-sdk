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

type CommitmentPlansUpdatePlanOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CommitmentPlansUpdatePlan ...
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansUpdatePlan(ctx context.Context, id CommitmentPlanId, input PatchResourceTagsAndSku) (result CommitmentPlansUpdatePlanOperationResponse, err error) {
	req, err := c.preparerForCommitmentPlansUpdatePlan(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansUpdatePlan", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCommitmentPlansUpdatePlan(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansUpdatePlan", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CommitmentPlansUpdatePlanThenPoll performs CommitmentPlansUpdatePlan then polls until it's completed
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansUpdatePlanThenPoll(ctx context.Context, id CommitmentPlanId, input PatchResourceTagsAndSku) error {
	result, err := c.CommitmentPlansUpdatePlan(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CommitmentPlansUpdatePlan: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CommitmentPlansUpdatePlan: %+v", err)
	}

	return nil
}

// preparerForCommitmentPlansUpdatePlan prepares the CommitmentPlansUpdatePlan request.
func (c CognitiveServicesCommitmentPlansClient) preparerForCommitmentPlansUpdatePlan(ctx context.Context, id CommitmentPlanId, input PatchResourceTagsAndSku) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCommitmentPlansUpdatePlan sends the CommitmentPlansUpdatePlan request. The method will close the
// http.Response Body if it receives an error.
func (c CognitiveServicesCommitmentPlansClient) senderForCommitmentPlansUpdatePlan(ctx context.Context, req *http.Request) (future CommitmentPlansUpdatePlanOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

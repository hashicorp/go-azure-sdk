package cognitiveservicescommitmentplans

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommitmentPlansGetPlanOperationResponse struct {
	HttpResponse *http.Response
	Model        *CommitmentPlan
}

// CommitmentPlansGetPlan ...
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansGetPlan(ctx context.Context, id CommitmentPlanId) (result CommitmentPlansGetPlanOperationResponse, err error) {
	req, err := c.preparerForCommitmentPlansGetPlan(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansGetPlan", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansGetPlan", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCommitmentPlansGetPlan(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cognitiveservicescommitmentplans.CognitiveServicesCommitmentPlansClient", "CommitmentPlansGetPlan", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCommitmentPlansGetPlan prepares the CommitmentPlansGetPlan request.
func (c CognitiveServicesCommitmentPlansClient) preparerForCommitmentPlansGetPlan(ctx context.Context, id CommitmentPlanId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCommitmentPlansGetPlan handles the response to the CommitmentPlansGetPlan request. The method always
// closes the http.Response Body.
func (c CognitiveServicesCommitmentPlansClient) responderForCommitmentPlansGetPlan(resp *http.Response) (result CommitmentPlansGetPlanOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

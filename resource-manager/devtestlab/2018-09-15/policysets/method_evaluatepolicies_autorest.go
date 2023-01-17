package policysets

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EvaluatePoliciesOperationResponse struct {
	HttpResponse *http.Response
	Model        *EvaluatePoliciesResponse
}

// EvaluatePolicies ...
func (c PolicySetsClient) EvaluatePolicies(ctx context.Context, id PolicySetId, input EvaluatePoliciesRequest) (result EvaluatePoliciesOperationResponse, err error) {
	req, err := c.preparerForEvaluatePolicies(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policysets.PolicySetsClient", "EvaluatePolicies", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "policysets.PolicySetsClient", "EvaluatePolicies", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForEvaluatePolicies(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policysets.PolicySetsClient", "EvaluatePolicies", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForEvaluatePolicies prepares the EvaluatePolicies request.
func (c PolicySetsClient) preparerForEvaluatePolicies(ctx context.Context, id PolicySetId, input EvaluatePoliciesRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/evaluatePolicies", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForEvaluatePolicies handles the response to the EvaluatePolicies request. The method always
// closes the http.Response Body.
func (c PolicySetsClient) responderForEvaluatePolicies(resp *http.Response) (result EvaluatePoliciesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

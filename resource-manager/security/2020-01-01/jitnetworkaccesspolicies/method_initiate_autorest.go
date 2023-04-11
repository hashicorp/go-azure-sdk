package jitnetworkaccesspolicies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InitiateOperationResponse struct {
	HttpResponse *http.Response
	Model        *JitNetworkAccessRequest
}

// Initiate ...
func (c JitNetworkAccessPoliciesClient) Initiate(ctx context.Context, id JitNetworkAccessPolicyId, input JitNetworkAccessPolicyInitiateRequest) (result InitiateOperationResponse, err error) {
	req, err := c.preparerForInitiate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "Initiate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "Initiate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForInitiate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jitnetworkaccesspolicies.JitNetworkAccessPoliciesClient", "Initiate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForInitiate prepares the Initiate request.
func (c JitNetworkAccessPoliciesClient) preparerForInitiate(ctx context.Context, id JitNetworkAccessPolicyId, input JitNetworkAccessPolicyInitiateRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/initiate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForInitiate handles the response to the Initiate request. The method always
// closes the http.Response Body.
func (c JitNetworkAccessPoliciesClient) responderForInitiate(resp *http.Response) (result InitiateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

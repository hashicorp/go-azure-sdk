package assessmentsmetadata

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// SubscriptionDelete ...
func (c AssessmentsMetadataClient) SubscriptionDelete(ctx context.Context, id ProviderAssessmentMetadataId) (result SubscriptionDeleteOperationResponse, err error) {
	req, err := c.preparerForSubscriptionDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSubscriptionDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSubscriptionDelete prepares the SubscriptionDelete request.
func (c AssessmentsMetadataClient) preparerForSubscriptionDelete(ctx context.Context, id ProviderAssessmentMetadataId) (*http.Request, error) {
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

// responderForSubscriptionDelete handles the response to the SubscriptionDelete request. The method always
// closes the http.Response Body.
func (c AssessmentsMetadataClient) responderForSubscriptionDelete(resp *http.Response) (result SubscriptionDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

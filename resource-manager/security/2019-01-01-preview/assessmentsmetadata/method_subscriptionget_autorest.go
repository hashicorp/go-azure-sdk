package assessmentsmetadata

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *SecurityAssessmentMetadata
}

// SubscriptionGet ...
func (c AssessmentsMetadataClient) SubscriptionGet(ctx context.Context, id ProviderAssessmentMetadataId) (result SubscriptionGetOperationResponse, err error) {
	req, err := c.preparerForSubscriptionGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSubscriptionGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSubscriptionGet prepares the SubscriptionGet request.
func (c AssessmentsMetadataClient) preparerForSubscriptionGet(ctx context.Context, id ProviderAssessmentMetadataId) (*http.Request, error) {
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

// responderForSubscriptionGet handles the response to the SubscriptionGet request. The method always
// closes the http.Response Body.
func (c AssessmentsMetadataClient) responderForSubscriptionGet(resp *http.Response) (result SubscriptionGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

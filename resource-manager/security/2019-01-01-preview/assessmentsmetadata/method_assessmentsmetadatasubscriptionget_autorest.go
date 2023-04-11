package assessmentsmetadata

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssessmentsMetadataSubscriptionGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *SecurityAssessmentMetadata
}

// AssessmentsMetadataSubscriptionGet ...
func (c AssessmentsMetadataClient) AssessmentsMetadataSubscriptionGet(ctx context.Context, id ProviderAssessmentMetadataId) (result AssessmentsMetadataSubscriptionGetOperationResponse, err error) {
	req, err := c.preparerForAssessmentsMetadataSubscriptionGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAssessmentsMetadataSubscriptionGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAssessmentsMetadataSubscriptionGet prepares the AssessmentsMetadataSubscriptionGet request.
func (c AssessmentsMetadataClient) preparerForAssessmentsMetadataSubscriptionGet(ctx context.Context, id ProviderAssessmentMetadataId) (*http.Request, error) {
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

// responderForAssessmentsMetadataSubscriptionGet handles the response to the AssessmentsMetadataSubscriptionGet request. The method always
// closes the http.Response Body.
func (c AssessmentsMetadataClient) responderForAssessmentsMetadataSubscriptionGet(resp *http.Response) (result AssessmentsMetadataSubscriptionGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

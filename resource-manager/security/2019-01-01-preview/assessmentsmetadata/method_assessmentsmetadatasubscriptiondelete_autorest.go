package assessmentsmetadata

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssessmentsMetadataSubscriptionDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// AssessmentsMetadataSubscriptionDelete ...
func (c AssessmentsMetadataClient) AssessmentsMetadataSubscriptionDelete(ctx context.Context, id ProviderAssessmentMetadataId) (result AssessmentsMetadataSubscriptionDeleteOperationResponse, err error) {
	req, err := c.preparerForAssessmentsMetadataSubscriptionDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAssessmentsMetadataSubscriptionDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAssessmentsMetadataSubscriptionDelete prepares the AssessmentsMetadataSubscriptionDelete request.
func (c AssessmentsMetadataClient) preparerForAssessmentsMetadataSubscriptionDelete(ctx context.Context, id ProviderAssessmentMetadataId) (*http.Request, error) {
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

// responderForAssessmentsMetadataSubscriptionDelete handles the response to the AssessmentsMetadataSubscriptionDelete request. The method always
// closes the http.Response Body.
func (c AssessmentsMetadataClient) responderForAssessmentsMetadataSubscriptionDelete(resp *http.Response) (result AssessmentsMetadataSubscriptionDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

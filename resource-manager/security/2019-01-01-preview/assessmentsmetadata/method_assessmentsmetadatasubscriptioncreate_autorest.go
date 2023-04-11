package assessmentsmetadata

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssessmentsMetadataSubscriptionCreateOperationResponse struct {
	HttpResponse *http.Response
	Model        *SecurityAssessmentMetadata
}

// AssessmentsMetadataSubscriptionCreate ...
func (c AssessmentsMetadataClient) AssessmentsMetadataSubscriptionCreate(ctx context.Context, id ProviderAssessmentMetadataId, input SecurityAssessmentMetadata) (result AssessmentsMetadataSubscriptionCreateOperationResponse, err error) {
	req, err := c.preparerForAssessmentsMetadataSubscriptionCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionCreate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionCreate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAssessmentsMetadataSubscriptionCreate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "AssessmentsMetadataSubscriptionCreate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAssessmentsMetadataSubscriptionCreate prepares the AssessmentsMetadataSubscriptionCreate request.
func (c AssessmentsMetadataClient) preparerForAssessmentsMetadataSubscriptionCreate(ctx context.Context, id ProviderAssessmentMetadataId, input SecurityAssessmentMetadata) (*http.Request, error) {
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

// responderForAssessmentsMetadataSubscriptionCreate handles the response to the AssessmentsMetadataSubscriptionCreate request. The method always
// closes the http.Response Body.
func (c AssessmentsMetadataClient) responderForAssessmentsMetadataSubscriptionCreate(resp *http.Response) (result AssessmentsMetadataSubscriptionCreateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

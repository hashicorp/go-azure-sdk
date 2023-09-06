package assessmentsmetadata

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionCreateOperationResponse struct {
	HttpResponse *http.Response
	Model        *SecurityAssessmentMetadata
}

// SubscriptionCreate ...
func (c AssessmentsMetadataClient) SubscriptionCreate(ctx context.Context, id ProviderAssessmentMetadataId, input SecurityAssessmentMetadata) (result SubscriptionCreateOperationResponse, err error) {
	req, err := c.preparerForSubscriptionCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionCreate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionCreate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSubscriptionCreate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "assessmentsmetadata.AssessmentsMetadataClient", "SubscriptionCreate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSubscriptionCreate prepares the SubscriptionCreate request.
func (c AssessmentsMetadataClient) preparerForSubscriptionCreate(ctx context.Context, id ProviderAssessmentMetadataId, input SecurityAssessmentMetadata) (*http.Request, error) {
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

// responderForSubscriptionCreate handles the response to the SubscriptionCreate request. The method always
// closes the http.Response Body.
func (c AssessmentsMetadataClient) responderForSubscriptionCreate(resp *http.Response) (result SubscriptionCreateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

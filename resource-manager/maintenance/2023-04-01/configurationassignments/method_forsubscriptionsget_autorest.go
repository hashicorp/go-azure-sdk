package configurationassignments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForSubscriptionsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConfigurationAssignment
}

// ForSubscriptionsGet ...
func (c ConfigurationAssignmentsClient) ForSubscriptionsGet(ctx context.Context, id ConfigurationAssignmentId) (result ForSubscriptionsGetOperationResponse, err error) {
	req, err := c.preparerForForSubscriptionsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForSubscriptionsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForSubscriptionsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForForSubscriptionsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForSubscriptionsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForForSubscriptionsGet prepares the ForSubscriptionsGet request.
func (c ConfigurationAssignmentsClient) preparerForForSubscriptionsGet(ctx context.Context, id ConfigurationAssignmentId) (*http.Request, error) {
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

// responderForForSubscriptionsGet handles the response to the ForSubscriptionsGet request. The method always
// closes the http.Response Body.
func (c ConfigurationAssignmentsClient) responderForForSubscriptionsGet(resp *http.Response) (result ForSubscriptionsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

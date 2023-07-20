package configurationassignments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForSubscriptionsDeleteOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConfigurationAssignment
}

// ForSubscriptionsDelete ...
func (c ConfigurationAssignmentsClient) ForSubscriptionsDelete(ctx context.Context, id ConfigurationAssignmentId) (result ForSubscriptionsDeleteOperationResponse, err error) {
	req, err := c.preparerForForSubscriptionsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForSubscriptionsDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForSubscriptionsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForForSubscriptionsDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForSubscriptionsDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForForSubscriptionsDelete prepares the ForSubscriptionsDelete request.
func (c ConfigurationAssignmentsClient) preparerForForSubscriptionsDelete(ctx context.Context, id ConfigurationAssignmentId) (*http.Request, error) {
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

// responderForForSubscriptionsDelete handles the response to the ForSubscriptionsDelete request. The method always
// closes the http.Response Body.
func (c ConfigurationAssignmentsClient) responderForForSubscriptionsDelete(resp *http.Response) (result ForSubscriptionsDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package configurationassignments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForSubscriptionsUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConfigurationAssignment
}

// ForSubscriptionsUpdate ...
func (c ConfigurationAssignmentsClient) ForSubscriptionsUpdate(ctx context.Context, id ConfigurationAssignmentId, input ConfigurationAssignment) (result ForSubscriptionsUpdateOperationResponse, err error) {
	req, err := c.preparerForForSubscriptionsUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForSubscriptionsUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForSubscriptionsUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForForSubscriptionsUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForSubscriptionsUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForForSubscriptionsUpdate prepares the ForSubscriptionsUpdate request.
func (c ConfigurationAssignmentsClient) preparerForForSubscriptionsUpdate(ctx context.Context, id ConfigurationAssignmentId, input ConfigurationAssignment) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForForSubscriptionsUpdate handles the response to the ForSubscriptionsUpdate request. The method always
// closes the http.Response Body.
func (c ConfigurationAssignmentsClient) responderForForSubscriptionsUpdate(resp *http.Response) (result ForSubscriptionsUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

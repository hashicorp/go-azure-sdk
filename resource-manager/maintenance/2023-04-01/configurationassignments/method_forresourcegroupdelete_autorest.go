package configurationassignments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForResourceGroupDeleteOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConfigurationAssignment
}

// ForResourceGroupDelete ...
func (c ConfigurationAssignmentsClient) ForResourceGroupDelete(ctx context.Context, id ProviderConfigurationAssignmentId) (result ForResourceGroupDeleteOperationResponse, err error) {
	req, err := c.preparerForForResourceGroupDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForResourceGroupDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForResourceGroupDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForForResourceGroupDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForResourceGroupDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForForResourceGroupDelete prepares the ForResourceGroupDelete request.
func (c ConfigurationAssignmentsClient) preparerForForResourceGroupDelete(ctx context.Context, id ProviderConfigurationAssignmentId) (*http.Request, error) {
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

// responderForForResourceGroupDelete handles the response to the ForResourceGroupDelete request. The method always
// closes the http.Response Body.
func (c ConfigurationAssignmentsClient) responderForForResourceGroupDelete(resp *http.Response) (result ForResourceGroupDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusNoContent, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

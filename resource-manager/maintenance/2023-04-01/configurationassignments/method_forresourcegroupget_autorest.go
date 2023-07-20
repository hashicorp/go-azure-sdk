package configurationassignments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForResourceGroupGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConfigurationAssignment
}

// ForResourceGroupGet ...
func (c ConfigurationAssignmentsClient) ForResourceGroupGet(ctx context.Context, id ProviderConfigurationAssignmentId) (result ForResourceGroupGetOperationResponse, err error) {
	req, err := c.preparerForForResourceGroupGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForResourceGroupGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForResourceGroupGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForForResourceGroupGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForResourceGroupGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForForResourceGroupGet prepares the ForResourceGroupGet request.
func (c ConfigurationAssignmentsClient) preparerForForResourceGroupGet(ctx context.Context, id ProviderConfigurationAssignmentId) (*http.Request, error) {
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

// responderForForResourceGroupGet handles the response to the ForResourceGroupGet request. The method always
// closes the http.Response Body.
func (c ConfigurationAssignmentsClient) responderForForResourceGroupGet(resp *http.Response) (result ForResourceGroupGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

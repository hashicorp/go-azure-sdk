package configurationassignments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForResourceGroupCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConfigurationAssignment
}

// ForResourceGroupCreateOrUpdate ...
func (c ConfigurationAssignmentsClient) ForResourceGroupCreateOrUpdate(ctx context.Context, id ProviderConfigurationAssignmentId, input ConfigurationAssignment) (result ForResourceGroupCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForForResourceGroupCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForResourceGroupCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForResourceGroupCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForForResourceGroupCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "configurationassignments.ConfigurationAssignmentsClient", "ForResourceGroupCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForForResourceGroupCreateOrUpdate prepares the ForResourceGroupCreateOrUpdate request.
func (c ConfigurationAssignmentsClient) preparerForForResourceGroupCreateOrUpdate(ctx context.Context, id ProviderConfigurationAssignmentId, input ConfigurationAssignment) (*http.Request, error) {
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

// responderForForResourceGroupCreateOrUpdate handles the response to the ForResourceGroupCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c ConfigurationAssignmentsClient) responderForForResourceGroupCreateOrUpdate(resp *http.Response) (result ForResourceGroupCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

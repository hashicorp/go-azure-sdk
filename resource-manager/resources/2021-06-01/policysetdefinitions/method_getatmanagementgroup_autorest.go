package policysetdefinitions

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAtManagementGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *PolicySetDefinition
}

// GetAtManagementGroup ...
func (c PolicySetDefinitionsClient) GetAtManagementGroup(ctx context.Context, id Providers2PolicySetDefinitionId) (result GetAtManagementGroupOperationResponse, err error) {
	req, err := c.preparerForGetAtManagementGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policysetdefinitions.PolicySetDefinitionsClient", "GetAtManagementGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "policysetdefinitions.PolicySetDefinitionsClient", "GetAtManagementGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAtManagementGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policysetdefinitions.PolicySetDefinitionsClient", "GetAtManagementGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAtManagementGroup prepares the GetAtManagementGroup request.
func (c PolicySetDefinitionsClient) preparerForGetAtManagementGroup(ctx context.Context, id Providers2PolicySetDefinitionId) (*http.Request, error) {
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

// responderForGetAtManagementGroup handles the response to the GetAtManagementGroup request. The method always
// closes the http.Response Body.
func (c PolicySetDefinitionsClient) responderForGetAtManagementGroup(resp *http.Response) (result GetAtManagementGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package policydefinitions

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateAtManagementGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *PolicyDefinition
}

// CreateOrUpdateAtManagementGroup ...
func (c PolicyDefinitionsClient) CreateOrUpdateAtManagementGroup(ctx context.Context, id Providers2PolicyDefinitionId, input PolicyDefinition) (result CreateOrUpdateAtManagementGroupOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateAtManagementGroup(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "CreateOrUpdateAtManagementGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "CreateOrUpdateAtManagementGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateAtManagementGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "CreateOrUpdateAtManagementGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateAtManagementGroup prepares the CreateOrUpdateAtManagementGroup request.
func (c PolicyDefinitionsClient) preparerForCreateOrUpdateAtManagementGroup(ctx context.Context, id Providers2PolicyDefinitionId, input PolicyDefinition) (*http.Request, error) {
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

// responderForCreateOrUpdateAtManagementGroup handles the response to the CreateOrUpdateAtManagementGroup request. The method always
// closes the http.Response Body.
func (c PolicyDefinitionsClient) responderForCreateOrUpdateAtManagementGroup(resp *http.Response) (result CreateOrUpdateAtManagementGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

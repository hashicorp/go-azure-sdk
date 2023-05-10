package policydefinitions

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteAtManagementGroupOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteAtManagementGroup ...
func (c PolicyDefinitionsClient) DeleteAtManagementGroup(ctx context.Context, id Providers2PolicyDefinitionId) (result DeleteAtManagementGroupOperationResponse, err error) {
	req, err := c.preparerForDeleteAtManagementGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "DeleteAtManagementGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "DeleteAtManagementGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteAtManagementGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "DeleteAtManagementGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteAtManagementGroup prepares the DeleteAtManagementGroup request.
func (c PolicyDefinitionsClient) preparerForDeleteAtManagementGroup(ctx context.Context, id Providers2PolicyDefinitionId) (*http.Request, error) {
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

// responderForDeleteAtManagementGroup handles the response to the DeleteAtManagementGroup request. The method always
// closes the http.Response Body.
func (c PolicyDefinitionsClient) responderForDeleteAtManagementGroup(resp *http.Response) (result DeleteAtManagementGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

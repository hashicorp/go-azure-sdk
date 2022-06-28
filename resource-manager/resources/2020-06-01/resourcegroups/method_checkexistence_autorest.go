package resourcegroups

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckExistenceOperationResponse struct {
	HttpResponse *http.Response
}

// CheckExistence ...
func (c ResourceGroupsClient) CheckExistence(ctx context.Context, id commonids.ResourceGroupId) (result CheckExistenceOperationResponse, err error) {
	req, err := c.preparerForCheckExistence(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegroups.ResourceGroupsClient", "CheckExistence", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegroups.ResourceGroupsClient", "CheckExistence", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCheckExistence(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcegroups.ResourceGroupsClient", "CheckExistence", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCheckExistence prepares the CheckExistence request.
func (c ResourceGroupsClient) preparerForCheckExistence(ctx context.Context, id commonids.ResourceGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsHead(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCheckExistence handles the response to the CheckExistence request. The method always
// closes the http.Response Body.
func (c ResourceGroupsClient) responderForCheckExistence(resp *http.Response) (result CheckExistenceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

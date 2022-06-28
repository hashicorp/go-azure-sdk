package resources

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckExistenceByIdOperationResponse struct {
	HttpResponse *http.Response
}

// CheckExistenceById ...
func (c ResourcesClient) CheckExistenceById(ctx context.Context, id commonids.ScopeId) (result CheckExistenceByIdOperationResponse, err error) {
	req, err := c.preparerForCheckExistenceById(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ResourcesClient", "CheckExistenceById", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ResourcesClient", "CheckExistenceById", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCheckExistenceById(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ResourcesClient", "CheckExistenceById", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCheckExistenceById prepares the CheckExistenceById request.
func (c ResourcesClient) preparerForCheckExistenceById(ctx context.Context, id commonids.ScopeId) (*http.Request, error) {
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

// responderForCheckExistenceById handles the response to the CheckExistenceById request. The method always
// closes the http.Response Body.
func (c ResourcesClient) responderForCheckExistenceById(resp *http.Response) (result CheckExistenceByIdOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

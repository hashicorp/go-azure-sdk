package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteFunctionSecretSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteFunctionSecretSlot ...
func (c WebAppsClient) DeleteFunctionSecretSlot(ctx context.Context, id FunctionKeyId) (result DeleteFunctionSecretSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteFunctionSecretSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteFunctionSecretSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteFunctionSecretSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteFunctionSecretSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteFunctionSecretSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteFunctionSecretSlot prepares the DeleteFunctionSecretSlot request.
func (c WebAppsClient) preparerForDeleteFunctionSecretSlot(ctx context.Context, id FunctionKeyId) (*http.Request, error) {
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

// responderForDeleteFunctionSecretSlot handles the response to the DeleteFunctionSecretSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteFunctionSecretSlot(resp *http.Response) (result DeleteFunctionSecretSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

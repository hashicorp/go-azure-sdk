package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteHostSecretSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteHostSecretSlot ...
func (c WebAppsClient) DeleteHostSecretSlot(ctx context.Context, id HostDefaultId) (result DeleteHostSecretSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteHostSecretSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHostSecretSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHostSecretSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteHostSecretSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteHostSecretSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteHostSecretSlot prepares the DeleteHostSecretSlot request.
func (c WebAppsClient) preparerForDeleteHostSecretSlot(ctx context.Context, id HostDefaultId) (*http.Request, error) {
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

// responderForDeleteHostSecretSlot handles the response to the DeleteHostSecretSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteHostSecretSlot(resp *http.Response) (result DeleteHostSecretSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

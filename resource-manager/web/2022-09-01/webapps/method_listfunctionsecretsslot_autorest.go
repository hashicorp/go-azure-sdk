package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListFunctionSecretsSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *FunctionSecrets
}

// ListFunctionSecretsSlot ...
func (c WebAppsClient) ListFunctionSecretsSlot(ctx context.Context, id SlotFunctionId) (result ListFunctionSecretsSlotOperationResponse, err error) {
	req, err := c.preparerForListFunctionSecretsSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctionSecretsSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctionSecretsSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListFunctionSecretsSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctionSecretsSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListFunctionSecretsSlot prepares the ListFunctionSecretsSlot request.
func (c WebAppsClient) preparerForListFunctionSecretsSlot(ctx context.Context, id SlotFunctionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listsecrets", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListFunctionSecretsSlot handles the response to the ListFunctionSecretsSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListFunctionSecretsSlot(resp *http.Response) (result ListFunctionSecretsSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

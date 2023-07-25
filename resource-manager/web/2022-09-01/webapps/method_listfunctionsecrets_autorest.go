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

type ListFunctionSecretsOperationResponse struct {
	HttpResponse *http.Response
	Model        *FunctionSecrets
}

// ListFunctionSecrets ...
func (c WebAppsClient) ListFunctionSecrets(ctx context.Context, id FunctionId) (result ListFunctionSecretsOperationResponse, err error) {
	req, err := c.preparerForListFunctionSecrets(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctionSecrets", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctionSecrets", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListFunctionSecrets(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctionSecrets", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListFunctionSecrets prepares the ListFunctionSecrets request.
func (c WebAppsClient) preparerForListFunctionSecrets(ctx context.Context, id FunctionId) (*http.Request, error) {
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

// responderForListFunctionSecrets handles the response to the ListFunctionSecrets request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListFunctionSecrets(resp *http.Response) (result ListFunctionSecretsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

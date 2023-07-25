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

type ListFunctionKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// ListFunctionKeys ...
func (c WebAppsClient) ListFunctionKeys(ctx context.Context, id FunctionId) (result ListFunctionKeysOperationResponse, err error) {
	req, err := c.preparerForListFunctionKeys(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctionKeys", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctionKeys", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListFunctionKeys(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListFunctionKeys", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListFunctionKeys prepares the ListFunctionKeys request.
func (c WebAppsClient) preparerForListFunctionKeys(ctx context.Context, id FunctionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listkeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListFunctionKeys handles the response to the ListFunctionKeys request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListFunctionKeys(resp *http.Response) (result ListFunctionKeysOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

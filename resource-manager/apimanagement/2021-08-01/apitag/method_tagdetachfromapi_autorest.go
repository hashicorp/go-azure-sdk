package apitag

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagDetachFromApiOperationResponse struct {
	HttpResponse *http.Response
}

// TagDetachFromApi ...
func (c ApiTagClient) TagDetachFromApi(ctx context.Context, id ApiTagId) (result TagDetachFromApiOperationResponse, err error) {
	req, err := c.preparerForTagDetachFromApi(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagDetachFromApi", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagDetachFromApi", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTagDetachFromApi(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagDetachFromApi", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTagDetachFromApi prepares the TagDetachFromApi request.
func (c ApiTagClient) preparerForTagDetachFromApi(ctx context.Context, id ApiTagId) (*http.Request, error) {
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

// responderForTagDetachFromApi handles the response to the TagDetachFromApi request. The method always
// closes the http.Response Body.
func (c ApiTagClient) responderForTagDetachFromApi(resp *http.Response) (result TagDetachFromApiOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

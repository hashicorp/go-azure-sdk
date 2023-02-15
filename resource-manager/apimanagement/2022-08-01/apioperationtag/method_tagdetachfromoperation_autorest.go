package apioperationtag

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagDetachFromOperationOperationResponse struct {
	HttpResponse *http.Response
}

// TagDetachFromOperation ...
func (c ApiOperationTagClient) TagDetachFromOperation(ctx context.Context, id OperationTagId) (result TagDetachFromOperationOperationResponse, err error) {
	req, err := c.preparerForTagDetachFromOperation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationtag.ApiOperationTagClient", "TagDetachFromOperation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationtag.ApiOperationTagClient", "TagDetachFromOperation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTagDetachFromOperation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationtag.ApiOperationTagClient", "TagDetachFromOperation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTagDetachFromOperation prepares the TagDetachFromOperation request.
func (c ApiOperationTagClient) preparerForTagDetachFromOperation(ctx context.Context, id OperationTagId) (*http.Request, error) {
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

// responderForTagDetachFromOperation handles the response to the TagDetachFromOperation request. The method always
// closes the http.Response Body.
func (c ApiOperationTagClient) responderForTagDetachFromOperation(resp *http.Response) (result TagDetachFromOperationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

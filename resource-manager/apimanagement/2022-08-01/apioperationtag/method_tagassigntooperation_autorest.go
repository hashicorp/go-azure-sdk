package apioperationtag

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagAssignToOperationOperationResponse struct {
	HttpResponse *http.Response
	Model        *TagContract
}

// TagAssignToOperation ...
func (c ApiOperationTagClient) TagAssignToOperation(ctx context.Context, id OperationTagId) (result TagAssignToOperationOperationResponse, err error) {
	req, err := c.preparerForTagAssignToOperation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationtag.ApiOperationTagClient", "TagAssignToOperation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationtag.ApiOperationTagClient", "TagAssignToOperation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTagAssignToOperation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationtag.ApiOperationTagClient", "TagAssignToOperation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTagAssignToOperation prepares the TagAssignToOperation request.
func (c ApiOperationTagClient) preparerForTagAssignToOperation(ctx context.Context, id OperationTagId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForTagAssignToOperation handles the response to the TagAssignToOperation request. The method always
// closes the http.Response Body.
func (c ApiOperationTagClient) responderForTagAssignToOperation(resp *http.Response) (result TagAssignToOperationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

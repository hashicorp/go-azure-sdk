package apioperationtag

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagGetByOperationOperationResponse struct {
	HttpResponse *http.Response
	Model        *TagContract
}

// TagGetByOperation ...
func (c ApiOperationTagClient) TagGetByOperation(ctx context.Context, id OperationTagId) (result TagGetByOperationOperationResponse, err error) {
	req, err := c.preparerForTagGetByOperation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationtag.ApiOperationTagClient", "TagGetByOperation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationtag.ApiOperationTagClient", "TagGetByOperation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTagGetByOperation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationtag.ApiOperationTagClient", "TagGetByOperation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTagGetByOperation prepares the TagGetByOperation request.
func (c ApiOperationTagClient) preparerForTagGetByOperation(ctx context.Context, id OperationTagId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForTagGetByOperation handles the response to the TagGetByOperation request. The method always
// closes the http.Response Body.
func (c ApiOperationTagClient) responderForTagGetByOperation(resp *http.Response) (result TagGetByOperationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

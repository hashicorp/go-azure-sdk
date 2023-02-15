package apioperationtag

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagGetEntityStateByOperationOperationResponse struct {
	HttpResponse *http.Response
}

// TagGetEntityStateByOperation ...
func (c ApiOperationTagClient) TagGetEntityStateByOperation(ctx context.Context, id OperationTagId) (result TagGetEntityStateByOperationOperationResponse, err error) {
	req, err := c.preparerForTagGetEntityStateByOperation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationtag.ApiOperationTagClient", "TagGetEntityStateByOperation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationtag.ApiOperationTagClient", "TagGetEntityStateByOperation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTagGetEntityStateByOperation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apioperationtag.ApiOperationTagClient", "TagGetEntityStateByOperation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTagGetEntityStateByOperation prepares the TagGetEntityStateByOperation request.
func (c ApiOperationTagClient) preparerForTagGetEntityStateByOperation(ctx context.Context, id OperationTagId) (*http.Request, error) {
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

// responderForTagGetEntityStateByOperation handles the response to the TagGetEntityStateByOperation request. The method always
// closes the http.Response Body.
func (c ApiOperationTagClient) responderForTagGetEntityStateByOperation(resp *http.Response) (result TagGetEntityStateByOperationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

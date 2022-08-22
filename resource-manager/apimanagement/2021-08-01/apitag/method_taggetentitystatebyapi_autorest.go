package apitag

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagGetEntityStateByApiOperationResponse struct {
	HttpResponse *http.Response
}

// TagGetEntityStateByApi ...
func (c ApiTagClient) TagGetEntityStateByApi(ctx context.Context, id TagId) (result TagGetEntityStateByApiOperationResponse, err error) {
	req, err := c.preparerForTagGetEntityStateByApi(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagGetEntityStateByApi", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagGetEntityStateByApi", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTagGetEntityStateByApi(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagGetEntityStateByApi", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTagGetEntityStateByApi prepares the TagGetEntityStateByApi request.
func (c ApiTagClient) preparerForTagGetEntityStateByApi(ctx context.Context, id TagId) (*http.Request, error) {
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

// responderForTagGetEntityStateByApi handles the response to the TagGetEntityStateByApi request. The method always
// closes the http.Response Body.
func (c ApiTagClient) responderForTagGetEntityStateByApi(resp *http.Response) (result TagGetEntityStateByApiOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

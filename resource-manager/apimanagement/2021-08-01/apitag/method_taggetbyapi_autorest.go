package apitag

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagGetByApiOperationResponse struct {
	HttpResponse *http.Response
	Model        *TagContract
}

// TagGetByApi ...
func (c ApiTagClient) TagGetByApi(ctx context.Context, id ApiTagId) (result TagGetByApiOperationResponse, err error) {
	req, err := c.preparerForTagGetByApi(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagGetByApi", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagGetByApi", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTagGetByApi(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagGetByApi", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTagGetByApi prepares the TagGetByApi request.
func (c ApiTagClient) preparerForTagGetByApi(ctx context.Context, id ApiTagId) (*http.Request, error) {
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

// responderForTagGetByApi handles the response to the TagGetByApi request. The method always
// closes the http.Response Body.
func (c ApiTagClient) responderForTagGetByApi(resp *http.Response) (result TagGetByApiOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

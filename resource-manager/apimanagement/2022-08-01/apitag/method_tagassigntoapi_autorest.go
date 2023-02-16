package apitag

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagAssignToApiOperationResponse struct {
	HttpResponse *http.Response
	Model        *TagContract
}

// TagAssignToApi ...
func (c ApiTagClient) TagAssignToApi(ctx context.Context, id TagId) (result TagAssignToApiOperationResponse, err error) {
	req, err := c.preparerForTagAssignToApi(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagAssignToApi", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagAssignToApi", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTagAssignToApi(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apitag.ApiTagClient", "TagAssignToApi", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTagAssignToApi prepares the TagAssignToApi request.
func (c ApiTagClient) preparerForTagAssignToApi(ctx context.Context, id TagId) (*http.Request, error) {
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

// responderForTagAssignToApi handles the response to the TagAssignToApi request. The method always
// closes the http.Response Body.
func (c ApiTagClient) responderForTagAssignToApi(resp *http.Response) (result TagAssignToApiOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

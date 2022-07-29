package media

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentKeyPoliciesDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// ContentKeyPoliciesDelete ...
func (c MediaClient) ContentKeyPoliciesDelete(ctx context.Context, id ContentKeyPoliciesId) (result ContentKeyPoliciesDeleteOperationResponse, err error) {
	req, err := c.preparerForContentKeyPoliciesDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.MediaClient", "ContentKeyPoliciesDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.MediaClient", "ContentKeyPoliciesDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContentKeyPoliciesDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.MediaClient", "ContentKeyPoliciesDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContentKeyPoliciesDelete prepares the ContentKeyPoliciesDelete request.
func (c MediaClient) preparerForContentKeyPoliciesDelete(ctx context.Context, id ContentKeyPoliciesId) (*http.Request, error) {
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

// responderForContentKeyPoliciesDelete handles the response to the ContentKeyPoliciesDelete request. The method always
// closes the http.Response Body.
func (c MediaClient) responderForContentKeyPoliciesDelete(resp *http.Response) (result ContentKeyPoliciesDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package media

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentKeyPoliciesCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ContentKeyPolicy
}

// ContentKeyPoliciesCreateOrUpdate ...
func (c MediaClient) ContentKeyPoliciesCreateOrUpdate(ctx context.Context, id ContentKeyPoliciesId, input ContentKeyPolicy) (result ContentKeyPoliciesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForContentKeyPoliciesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.MediaClient", "ContentKeyPoliciesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.MediaClient", "ContentKeyPoliciesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContentKeyPoliciesCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.MediaClient", "ContentKeyPoliciesCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContentKeyPoliciesCreateOrUpdate prepares the ContentKeyPoliciesCreateOrUpdate request.
func (c MediaClient) preparerForContentKeyPoliciesCreateOrUpdate(ctx context.Context, id ContentKeyPoliciesId, input ContentKeyPolicy) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForContentKeyPoliciesCreateOrUpdate handles the response to the ContentKeyPoliciesCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c MediaClient) responderForContentKeyPoliciesCreateOrUpdate(resp *http.Response) (result ContentKeyPoliciesCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

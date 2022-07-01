package media

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StreamingPoliciesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *StreamingPolicy
}

// StreamingPoliciesGet ...
func (c MediaClient) StreamingPoliciesGet(ctx context.Context, id StreamingPoliciesId) (result StreamingPoliciesGetOperationResponse, err error) {
	req, err := c.preparerForStreamingPoliciesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.MediaClient", "StreamingPoliciesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.MediaClient", "StreamingPoliciesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStreamingPoliciesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "media.MediaClient", "StreamingPoliciesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStreamingPoliciesGet prepares the StreamingPoliciesGet request.
func (c MediaClient) preparerForStreamingPoliciesGet(ctx context.Context, id StreamingPoliciesId) (*http.Request, error) {
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

// responderForStreamingPoliciesGet handles the response to the StreamingPoliciesGet request. The method always
// closes the http.Response Body.
func (c MediaClient) responderForStreamingPoliciesGet(resp *http.Response) (result StreamingPoliciesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

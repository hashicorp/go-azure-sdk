package privatelinkresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetByGroupIdOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateLinkResource
}

// GetByGroupId ...
func (c PrivateLinkResourceClient) GetByGroupId(ctx context.Context, id PrivateLinkResourceId) (result GetByGroupIdOperationResponse, err error) {
	req, err := c.preparerForGetByGroupId(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkresource.PrivateLinkResourceClient", "GetByGroupId", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkresource.PrivateLinkResourceClient", "GetByGroupId", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetByGroupId(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkresource.PrivateLinkResourceClient", "GetByGroupId", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetByGroupId prepares the GetByGroupId request.
func (c PrivateLinkResourceClient) preparerForGetByGroupId(ctx context.Context, id PrivateLinkResourceId) (*http.Request, error) {
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

// responderForGetByGroupId handles the response to the GetByGroupId request. The method always
// closes the http.Response Body.
func (c PrivateLinkResourceClient) responderForGetByGroupId(resp *http.Response) (result GetByGroupIdOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

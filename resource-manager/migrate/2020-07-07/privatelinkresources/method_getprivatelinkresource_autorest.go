package privatelinkresources

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPrivateLinkResourceOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateLinkResource
}

// GetPrivateLinkResource ...
func (c PrivateLinkResourcesClient) GetPrivateLinkResource(ctx context.Context, id PrivateLinkResourceId) (result GetPrivateLinkResourceOperationResponse, err error) {
	req, err := c.preparerForGetPrivateLinkResource(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkresources.PrivateLinkResourcesClient", "GetPrivateLinkResource", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkresources.PrivateLinkResourcesClient", "GetPrivateLinkResource", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetPrivateLinkResource(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privatelinkresources.PrivateLinkResourcesClient", "GetPrivateLinkResource", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetPrivateLinkResource prepares the GetPrivateLinkResource request.
func (c PrivateLinkResourcesClient) preparerForGetPrivateLinkResource(ctx context.Context, id PrivateLinkResourceId) (*http.Request, error) {
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

// responderForGetPrivateLinkResource handles the response to the GetPrivateLinkResource request. The method always
// closes the http.Response Body.
func (c PrivateLinkResourcesClient) responderForGetPrivateLinkResource(resp *http.Response) (result GetPrivateLinkResourceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

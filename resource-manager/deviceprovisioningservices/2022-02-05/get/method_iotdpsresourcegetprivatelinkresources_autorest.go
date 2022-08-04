package get

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotDpsResourceGetPrivateLinkResourcesOperationResponse struct {
	HttpResponse *http.Response
	Model        *GroupIdInformation
}

// IotDpsResourceGetPrivateLinkResources ...
func (c GETClient) IotDpsResourceGetPrivateLinkResources(ctx context.Context, id PrivateLinkResourceId) (result IotDpsResourceGetPrivateLinkResourcesOperationResponse, err error) {
	req, err := c.preparerForIotDpsResourceGetPrivateLinkResources(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceGetPrivateLinkResources", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceGetPrivateLinkResources", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIotDpsResourceGetPrivateLinkResources(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "IotDpsResourceGetPrivateLinkResources", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIotDpsResourceGetPrivateLinkResources prepares the IotDpsResourceGetPrivateLinkResources request.
func (c GETClient) preparerForIotDpsResourceGetPrivateLinkResources(ctx context.Context, id PrivateLinkResourceId) (*http.Request, error) {
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

// responderForIotDpsResourceGetPrivateLinkResources handles the response to the IotDpsResourceGetPrivateLinkResources request. The method always
// closes the http.Response Body.
func (c GETClient) responderForIotDpsResourceGetPrivateLinkResources(resp *http.Response) (result IotDpsResourceGetPrivateLinkResourcesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

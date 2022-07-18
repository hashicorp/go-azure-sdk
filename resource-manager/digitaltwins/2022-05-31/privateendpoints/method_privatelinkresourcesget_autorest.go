package privateendpoints

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkResourcesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *GroupIdInformation
}

// PrivateLinkResourcesGet ...
func (c PrivateEndpointsClient) PrivateLinkResourcesGet(ctx context.Context, id ResourceId) (result PrivateLinkResourcesGetOperationResponse, err error) {
	req, err := c.preparerForPrivateLinkResourcesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpoints.PrivateEndpointsClient", "PrivateLinkResourcesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpoints.PrivateEndpointsClient", "PrivateLinkResourcesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateLinkResourcesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpoints.PrivateEndpointsClient", "PrivateLinkResourcesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateLinkResourcesGet prepares the PrivateLinkResourcesGet request.
func (c PrivateEndpointsClient) preparerForPrivateLinkResourcesGet(ctx context.Context, id ResourceId) (*http.Request, error) {
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

// responderForPrivateLinkResourcesGet handles the response to the PrivateLinkResourcesGet request. The method always
// closes the http.Response Body.
func (c PrivateEndpointsClient) responderForPrivateLinkResourcesGet(resp *http.Response) (result PrivateLinkResourcesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

package cloudendpointresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudEndpointsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *CloudEndpoint
}

// CloudEndpointsGet ...
func (c CloudEndpointResourceClient) CloudEndpointsGet(ctx context.Context, id CloudEndpointId) (result CloudEndpointsGetOperationResponse, err error) {
	req, err := c.preparerForCloudEndpointsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudendpointresource.CloudEndpointResourceClient", "CloudEndpointsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudendpointresource.CloudEndpointResourceClient", "CloudEndpointsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCloudEndpointsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudendpointresource.CloudEndpointResourceClient", "CloudEndpointsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCloudEndpointsGet prepares the CloudEndpointsGet request.
func (c CloudEndpointResourceClient) preparerForCloudEndpointsGet(ctx context.Context, id CloudEndpointId) (*http.Request, error) {
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

// responderForCloudEndpointsGet handles the response to the CloudEndpointsGet request. The method always
// closes the http.Response Body.
func (c CloudEndpointResourceClient) responderForCloudEndpointsGet(resp *http.Response) (result CloudEndpointsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package cloudendpointresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudEndpointsListBySyncGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *CloudEndpointArray
}

// CloudEndpointsListBySyncGroup ...
func (c CloudEndpointResourceClient) CloudEndpointsListBySyncGroup(ctx context.Context, id SyncGroupId) (result CloudEndpointsListBySyncGroupOperationResponse, err error) {
	req, err := c.preparerForCloudEndpointsListBySyncGroup(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudendpointresource.CloudEndpointResourceClient", "CloudEndpointsListBySyncGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudendpointresource.CloudEndpointResourceClient", "CloudEndpointsListBySyncGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCloudEndpointsListBySyncGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cloudendpointresource.CloudEndpointResourceClient", "CloudEndpointsListBySyncGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCloudEndpointsListBySyncGroup prepares the CloudEndpointsListBySyncGroup request.
func (c CloudEndpointResourceClient) preparerForCloudEndpointsListBySyncGroup(ctx context.Context, id SyncGroupId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/cloudEndpoints", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCloudEndpointsListBySyncGroup handles the response to the CloudEndpointsListBySyncGroup request. The method always
// closes the http.Response Body.
func (c CloudEndpointResourceClient) responderForCloudEndpointsListBySyncGroup(resp *http.Response) (result CloudEndpointsListBySyncGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

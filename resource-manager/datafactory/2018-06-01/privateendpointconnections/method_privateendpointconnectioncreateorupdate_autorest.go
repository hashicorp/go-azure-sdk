package privateendpointconnections

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionResource
}

type PrivateEndpointConnectionCreateOrUpdateOperationOptions struct {
	IfMatch *string
}

func DefaultPrivateEndpointConnectionCreateOrUpdateOperationOptions() PrivateEndpointConnectionCreateOrUpdateOperationOptions {
	return PrivateEndpointConnectionCreateOrUpdateOperationOptions{}
}

func (o PrivateEndpointConnectionCreateOrUpdateOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IfMatch != nil {
		out["If-Match"] = *o.IfMatch
	}

	return out
}

func (o PrivateEndpointConnectionCreateOrUpdateOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// PrivateEndpointConnectionCreateOrUpdate ...
func (c PrivateEndpointConnectionsClient) PrivateEndpointConnectionCreateOrUpdate(ctx context.Context, id PrivateEndpointConnectionId, input PrivateLinkConnectionApprovalRequestResource, options PrivateEndpointConnectionCreateOrUpdateOperationOptions) (result PrivateEndpointConnectionCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForPrivateEndpointConnectionCreateOrUpdate(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPrivateEndpointConnectionCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "PrivateEndpointConnectionCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPrivateEndpointConnectionCreateOrUpdate prepares the PrivateEndpointConnectionCreateOrUpdate request.
func (c PrivateEndpointConnectionsClient) preparerForPrivateEndpointConnectionCreateOrUpdate(ctx context.Context, id PrivateEndpointConnectionId, input PrivateLinkConnectionApprovalRequestResource, options PrivateEndpointConnectionCreateOrUpdateOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPrivateEndpointConnectionCreateOrUpdate handles the response to the PrivateEndpointConnectionCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionsClient) responderForPrivateEndpointConnectionCreateOrUpdate(resp *http.Response) (result PrivateEndpointConnectionCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

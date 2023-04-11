package privateendpointconnections

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDatabaseAccountOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateEndpointConnectionListResult
}

// ListByDatabaseAccount ...
func (c PrivateEndpointConnectionsClient) ListByDatabaseAccount(ctx context.Context, id DatabaseAccountId) (result ListByDatabaseAccountOperationResponse, err error) {
	req, err := c.preparerForListByDatabaseAccount(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "ListByDatabaseAccount", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "ListByDatabaseAccount", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByDatabaseAccount(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "privateendpointconnections.PrivateEndpointConnectionsClient", "ListByDatabaseAccount", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByDatabaseAccount prepares the ListByDatabaseAccount request.
func (c PrivateEndpointConnectionsClient) preparerForListByDatabaseAccount(ctx context.Context, id DatabaseAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateEndpointConnections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByDatabaseAccount handles the response to the ListByDatabaseAccount request. The method always
// closes the http.Response Body.
func (c PrivateEndpointConnectionsClient) responderForListByDatabaseAccount(resp *http.Response) (result ListByDatabaseAccountOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

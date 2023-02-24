package dataconnectorsdisconnect

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorsDisconnectOperationResponse struct {
	HttpResponse *http.Response
}

// DataConnectorsDisconnect ...
func (c DataConnectorsDisconnectClient) DataConnectorsDisconnect(ctx context.Context, id DataConnectorId) (result DataConnectorsDisconnectOperationResponse, err error) {
	req, err := c.preparerForDataConnectorsDisconnect(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataconnectorsdisconnect.DataConnectorsDisconnectClient", "DataConnectorsDisconnect", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataconnectorsdisconnect.DataConnectorsDisconnectClient", "DataConnectorsDisconnect", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDataConnectorsDisconnect(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataconnectorsdisconnect.DataConnectorsDisconnectClient", "DataConnectorsDisconnect", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDataConnectorsDisconnect prepares the DataConnectorsDisconnect request.
func (c DataConnectorsDisconnectClient) preparerForDataConnectorsDisconnect(ctx context.Context, id DataConnectorId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/disconnect", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDataConnectorsDisconnect handles the response to the DataConnectorsDisconnect request. The method always
// closes the http.Response Body.
func (c DataConnectorsDisconnectClient) responderForDataConnectorsDisconnect(resp *http.Response) (result DataConnectorsDisconnectOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

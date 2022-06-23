package dataconnectorsconnect

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorsConnectOperationResponse struct {
	HttpResponse *http.Response
}

// DataConnectorsConnect ...
func (c DataConnectorsConnectClient) DataConnectorsConnect(ctx context.Context, id DataConnectorId, input DataConnectorConnectBody) (result DataConnectorsConnectOperationResponse, err error) {
	req, err := c.preparerForDataConnectorsConnect(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataconnectorsconnect.DataConnectorsConnectClient", "DataConnectorsConnect", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataconnectorsconnect.DataConnectorsConnectClient", "DataConnectorsConnect", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDataConnectorsConnect(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataconnectorsconnect.DataConnectorsConnectClient", "DataConnectorsConnect", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDataConnectorsConnect prepares the DataConnectorsConnect request.
func (c DataConnectorsConnectClient) preparerForDataConnectorsConnect(ctx context.Context, id DataConnectorId, input DataConnectorConnectBody) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/connect", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDataConnectorsConnect handles the response to the DataConnectorsConnect request. The method always
// closes the http.Response Body.
func (c DataConnectorsConnectClient) responderForDataConnectorsConnect(resp *http.Response) (result DataConnectorsConnectOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

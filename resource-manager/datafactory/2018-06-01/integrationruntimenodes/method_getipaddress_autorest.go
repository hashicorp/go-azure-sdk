package integrationruntimenodes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetIPAddressOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeNodeIPAddress
}

// GetIPAddress ...
func (c IntegrationRuntimeNodesClient) GetIPAddress(ctx context.Context, id NodeId) (result GetIPAddressOperationResponse, err error) {
	req, err := c.preparerForGetIPAddress(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimenodes.IntegrationRuntimeNodesClient", "GetIPAddress", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimenodes.IntegrationRuntimeNodesClient", "GetIPAddress", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetIPAddress(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimenodes.IntegrationRuntimeNodesClient", "GetIPAddress", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetIPAddress prepares the GetIPAddress request.
func (c IntegrationRuntimeNodesClient) preparerForGetIPAddress(ctx context.Context, id NodeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/ipAddress", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetIPAddress handles the response to the GetIPAddress request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimeNodesClient) responderForGetIPAddress(resp *http.Response) (result GetIPAddressOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

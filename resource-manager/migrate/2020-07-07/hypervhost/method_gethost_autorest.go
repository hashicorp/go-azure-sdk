package hypervhost

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetHostOperationResponse struct {
	HttpResponse *http.Response
	Model        *HyperVHost
}

// GetHost ...
func (c HyperVHostClient) GetHost(ctx context.Context, id HostId) (result GetHostOperationResponse, err error) {
	req, err := c.preparerForGetHost(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervhost.HyperVHostClient", "GetHost", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervhost.HyperVHostClient", "GetHost", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetHost(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervhost.HyperVHostClient", "GetHost", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetHost prepares the GetHost request.
func (c HyperVHostClient) preparerForGetHost(ctx context.Context, id HostId) (*http.Request, error) {
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

// responderForGetHost handles the response to the GetHost request. The method always
// closes the http.Response Body.
func (c HyperVHostClient) responderForGetHost(resp *http.Response) (result GetHostOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

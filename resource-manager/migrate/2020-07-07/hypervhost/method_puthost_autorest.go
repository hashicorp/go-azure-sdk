package hypervhost

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PutHostOperationResponse struct {
	HttpResponse *http.Response
}

// PutHost ...
func (c HyperVHostClient) PutHost(ctx context.Context, id HostId, input HyperVHost) (result PutHostOperationResponse, err error) {
	req, err := c.preparerForPutHost(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervhost.HyperVHostClient", "PutHost", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervhost.HyperVHostClient", "PutHost", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPutHost(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervhost.HyperVHostClient", "PutHost", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPutHost prepares the PutHost request.
func (c HyperVHostClient) preparerForPutHost(ctx context.Context, id HostId, input HyperVHost) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPutHost handles the response to the PutHost request. The method always
// closes the http.Response Body.
func (c HyperVHostClient) responderForPutHost(resp *http.Response) (result PutHostOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

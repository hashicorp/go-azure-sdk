package machinenetworkprofile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkProfileGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *NetworkProfile
}

// NetworkProfileGet ...
func (c MachineNetworkProfileClient) NetworkProfileGet(ctx context.Context, id MachineId) (result NetworkProfileGetOperationResponse, err error) {
	req, err := c.preparerForNetworkProfileGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinenetworkprofile.MachineNetworkProfileClient", "NetworkProfileGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinenetworkprofile.MachineNetworkProfileClient", "NetworkProfileGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNetworkProfileGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "machinenetworkprofile.MachineNetworkProfileClient", "NetworkProfileGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNetworkProfileGet prepares the NetworkProfileGet request.
func (c MachineNetworkProfileClient) preparerForNetworkProfileGet(ctx context.Context, id MachineId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/networkProfile", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForNetworkProfileGet handles the response to the NetworkProfileGet request. The method always
// closes the http.Response Body.
func (c MachineNetworkProfileClient) responderForNetworkProfileGet(resp *http.Response) (result NetworkProfileGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

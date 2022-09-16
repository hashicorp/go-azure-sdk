package integrationruntime

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeIPAddressGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeNodeIPAddress
}

// NodeIPAddressGet ...
func (c IntegrationRuntimeClient) NodeIPAddressGet(ctx context.Context, id NodeId) (result NodeIPAddressGetOperationResponse, err error) {
	req, err := c.preparerForNodeIPAddressGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "NodeIPAddressGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "NodeIPAddressGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNodeIPAddressGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "NodeIPAddressGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNodeIPAddressGet prepares the NodeIPAddressGet request.
func (c IntegrationRuntimeClient) preparerForNodeIPAddressGet(ctx context.Context, id NodeId) (*http.Request, error) {
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

// responderForNodeIPAddressGet handles the response to the NodeIPAddressGet request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimeClient) responderForNodeIPAddressGet(resp *http.Response) (result NodeIPAddressGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

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

type ConnectionInfosGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeConnectionInfo
}

// ConnectionInfosGet ...
func (c IntegrationRuntimeClient) ConnectionInfosGet(ctx context.Context, id IntegrationRuntimeId) (result ConnectionInfosGetOperationResponse, err error) {
	req, err := c.preparerForConnectionInfosGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "ConnectionInfosGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "ConnectionInfosGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForConnectionInfosGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "ConnectionInfosGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForConnectionInfosGet prepares the ConnectionInfosGet request.
func (c IntegrationRuntimeClient) preparerForConnectionInfosGet(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getConnectionInfo", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForConnectionInfosGet handles the response to the ConnectionInfosGet request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimeClient) responderForConnectionInfosGet(resp *http.Response) (result ConnectionInfosGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

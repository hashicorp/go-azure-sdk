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

type StatusGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeStatusResponse
}

// StatusGet ...
func (c IntegrationRuntimeClient) StatusGet(ctx context.Context, id IntegrationRuntimeId) (result StatusGetOperationResponse, err error) {
	req, err := c.preparerForStatusGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "StatusGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "StatusGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStatusGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "StatusGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStatusGet prepares the StatusGet request.
func (c IntegrationRuntimeClient) preparerForStatusGet(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getStatus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStatusGet handles the response to the StatusGet request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimeClient) responderForStatusGet(resp *http.Response) (result StatusGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

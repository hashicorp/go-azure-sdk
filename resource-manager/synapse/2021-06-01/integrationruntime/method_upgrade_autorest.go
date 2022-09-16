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

type UpgradeOperationResponse struct {
	HttpResponse *http.Response
}

// Upgrade ...
func (c IntegrationRuntimeClient) Upgrade(ctx context.Context, id IntegrationRuntimeId) (result UpgradeOperationResponse, err error) {
	req, err := c.preparerForUpgrade(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "Upgrade", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "Upgrade", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpgrade(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "Upgrade", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpgrade prepares the Upgrade request.
func (c IntegrationRuntimeClient) preparerForUpgrade(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/upgrade", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpgrade handles the response to the Upgrade request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimeClient) responderForUpgrade(resp *http.Response) (result UpgradeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

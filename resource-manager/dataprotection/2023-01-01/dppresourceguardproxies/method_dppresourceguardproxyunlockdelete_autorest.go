package dppresourceguardproxies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DppResourceGuardProxyUnlockDeleteOperationResponse struct {
	HttpResponse *http.Response
	Model        *UnlockDeleteResponse
}

// DppResourceGuardProxyUnlockDelete ...
func (c DppResourceGuardProxiesClient) DppResourceGuardProxyUnlockDelete(ctx context.Context, id BackupResourceGuardProxyId, input UnlockDeleteRequest) (result DppResourceGuardProxyUnlockDeleteOperationResponse, err error) {
	req, err := c.preparerForDppResourceGuardProxyUnlockDelete(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dppresourceguardproxies.DppResourceGuardProxiesClient", "DppResourceGuardProxyUnlockDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dppresourceguardproxies.DppResourceGuardProxiesClient", "DppResourceGuardProxyUnlockDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDppResourceGuardProxyUnlockDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dppresourceguardproxies.DppResourceGuardProxiesClient", "DppResourceGuardProxyUnlockDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDppResourceGuardProxyUnlockDelete prepares the DppResourceGuardProxyUnlockDelete request.
func (c DppResourceGuardProxiesClient) preparerForDppResourceGuardProxyUnlockDelete(ctx context.Context, id BackupResourceGuardProxyId, input UnlockDeleteRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/unlockDelete", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDppResourceGuardProxyUnlockDelete handles the response to the DppResourceGuardProxyUnlockDelete request. The method always
// closes the http.Response Body.
func (c DppResourceGuardProxiesClient) responderForDppResourceGuardProxyUnlockDelete(resp *http.Response) (result DppResourceGuardProxyUnlockDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

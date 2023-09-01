package mhsmprivatelinkresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByMHSMResourceOperationResponse struct {
	HttpResponse *http.Response
	Model        *MHSMPrivateLinkResourceListResult
}

// ListByMHSMResource ...
func (c MHSMPrivateLinkResourcesClient) ListByMHSMResource(ctx context.Context, id ManagedHSMId) (result ListByMHSMResourceOperationResponse, err error) {
	req, err := c.preparerForListByMHSMResource(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mhsmprivatelinkresources.MHSMPrivateLinkResourcesClient", "ListByMHSMResource", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "mhsmprivatelinkresources.MHSMPrivateLinkResourcesClient", "ListByMHSMResource", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByMHSMResource(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "mhsmprivatelinkresources.MHSMPrivateLinkResourcesClient", "ListByMHSMResource", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByMHSMResource prepares the ListByMHSMResource request.
func (c MHSMPrivateLinkResourcesClient) preparerForListByMHSMResource(ctx context.Context, id ManagedHSMId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateLinkResources", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByMHSMResource handles the response to the ListByMHSMResource request. The method always
// closes the http.Response Body.
func (c MHSMPrivateLinkResourcesClient) responderForListByMHSMResource(resp *http.Response) (result ListByMHSMResourceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

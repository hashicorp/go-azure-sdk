package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListHostKeysSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *HostKeys
}

// ListHostKeysSlot ...
func (c WebAppsClient) ListHostKeysSlot(ctx context.Context, id SlotId) (result ListHostKeysSlotOperationResponse, err error) {
	req, err := c.preparerForListHostKeysSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostKeysSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostKeysSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListHostKeysSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListHostKeysSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListHostKeysSlot prepares the ListHostKeysSlot request.
func (c WebAppsClient) preparerForListHostKeysSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/host/default/listkeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListHostKeysSlot handles the response to the ListHostKeysSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListHostKeysSlot(resp *http.Response) (result ListHostKeysSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

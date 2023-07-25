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

type ListSyncFunctionTriggersSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *FunctionSecrets
}

// ListSyncFunctionTriggersSlot ...
func (c WebAppsClient) ListSyncFunctionTriggersSlot(ctx context.Context, id SlotId) (result ListSyncFunctionTriggersSlotOperationResponse, err error) {
	req, err := c.preparerForListSyncFunctionTriggersSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSyncFunctionTriggersSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSyncFunctionTriggersSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListSyncFunctionTriggersSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListSyncFunctionTriggersSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListSyncFunctionTriggersSlot prepares the ListSyncFunctionTriggersSlot request.
func (c WebAppsClient) preparerForListSyncFunctionTriggersSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listsyncfunctiontriggerstatus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListSyncFunctionTriggersSlot handles the response to the ListSyncFunctionTriggersSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListSyncFunctionTriggersSlot(resp *http.Response) (result ListSyncFunctionTriggersSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

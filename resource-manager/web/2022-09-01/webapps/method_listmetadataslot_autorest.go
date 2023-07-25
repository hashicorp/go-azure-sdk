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

type ListMetadataSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// ListMetadataSlot ...
func (c WebAppsClient) ListMetadataSlot(ctx context.Context, id SlotId) (result ListMetadataSlotOperationResponse, err error) {
	req, err := c.preparerForListMetadataSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListMetadataSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListMetadataSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListMetadataSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListMetadataSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListMetadataSlot prepares the ListMetadataSlot request.
func (c WebAppsClient) preparerForListMetadataSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/metadata/list", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListMetadataSlot handles the response to the ListMetadataSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListMetadataSlot(resp *http.Response) (result ListMetadataSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

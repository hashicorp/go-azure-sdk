package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StoragesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *StorageResource
}

// StoragesGet ...
func (c AppPlatformClient) StoragesGet(ctx context.Context, id StorageId) (result StoragesGetOperationResponse, err error) {
	req, err := c.preparerForStoragesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "StoragesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "StoragesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStoragesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "StoragesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStoragesGet prepares the StoragesGet request.
func (c AppPlatformClient) preparerForStoragesGet(ctx context.Context, id StorageId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStoragesGet handles the response to the StoragesGet request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForStoragesGet(resp *http.Response) (result StoragesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

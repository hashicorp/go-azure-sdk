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

type UpdateMetadataOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// UpdateMetadata ...
func (c WebAppsClient) UpdateMetadata(ctx context.Context, id SiteId, input StringDictionary) (result UpdateMetadataOperationResponse, err error) {
	req, err := c.preparerForUpdateMetadata(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateMetadata", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateMetadata", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateMetadata(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateMetadata", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateMetadata prepares the UpdateMetadata request.
func (c WebAppsClient) preparerForUpdateMetadata(ctx context.Context, id SiteId, input StringDictionary) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/metadata", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateMetadata handles the response to the UpdateMetadata request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateMetadata(resp *http.Response) (result UpdateMetadataOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

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

type ListMetadataOperationResponse struct {
	HttpResponse *http.Response
	Model        *StringDictionary
}

// ListMetadata ...
func (c WebAppsClient) ListMetadata(ctx context.Context, id SiteId) (result ListMetadataOperationResponse, err error) {
	req, err := c.preparerForListMetadata(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListMetadata", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListMetadata", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListMetadata(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListMetadata", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListMetadata prepares the ListMetadata request.
func (c WebAppsClient) preparerForListMetadata(ctx context.Context, id SiteId) (*http.Request, error) {
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

// responderForListMetadata handles the response to the ListMetadata request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListMetadata(resp *http.Response) (result ListMetadataOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

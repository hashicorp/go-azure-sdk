package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnlinkBackendOperationResponse struct {
	HttpResponse *http.Response
}

type UnlinkBackendOperationOptions struct {
	IsCleaningAuthConfig *bool
}

func DefaultUnlinkBackendOperationOptions() UnlinkBackendOperationOptions {
	return UnlinkBackendOperationOptions{}
}

func (o UnlinkBackendOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o UnlinkBackendOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IsCleaningAuthConfig != nil {
		out["isCleaningAuthConfig"] = *o.IsCleaningAuthConfig
	}

	return out
}

// UnlinkBackend ...
func (c StaticSitesClient) UnlinkBackend(ctx context.Context, id LinkedBackendId, options UnlinkBackendOperationOptions) (result UnlinkBackendOperationResponse, err error) {
	req, err := c.preparerForUnlinkBackend(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UnlinkBackend", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UnlinkBackend", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUnlinkBackend(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UnlinkBackend", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUnlinkBackend prepares the UnlinkBackend request.
func (c StaticSitesClient) preparerForUnlinkBackend(ctx context.Context, id LinkedBackendId, options UnlinkBackendOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUnlinkBackend handles the response to the UnlinkBackend request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForUnlinkBackend(resp *http.Response) (result UnlinkBackendOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

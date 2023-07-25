package staticsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnlinkBackendFromBuildOperationResponse struct {
	HttpResponse *http.Response
}

type UnlinkBackendFromBuildOperationOptions struct {
	IsCleaningAuthConfig *bool
}

func DefaultUnlinkBackendFromBuildOperationOptions() UnlinkBackendFromBuildOperationOptions {
	return UnlinkBackendFromBuildOperationOptions{}
}

func (o UnlinkBackendFromBuildOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o UnlinkBackendFromBuildOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IsCleaningAuthConfig != nil {
		out["isCleaningAuthConfig"] = *o.IsCleaningAuthConfig
	}

	return out
}

// UnlinkBackendFromBuild ...
func (c StaticSitesClient) UnlinkBackendFromBuild(ctx context.Context, id BuildLinkedBackendId, options UnlinkBackendFromBuildOperationOptions) (result UnlinkBackendFromBuildOperationResponse, err error) {
	req, err := c.preparerForUnlinkBackendFromBuild(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UnlinkBackendFromBuild", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UnlinkBackendFromBuild", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUnlinkBackendFromBuild(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "UnlinkBackendFromBuild", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUnlinkBackendFromBuild prepares the UnlinkBackendFromBuild request.
func (c StaticSitesClient) preparerForUnlinkBackendFromBuild(ctx context.Context, id BuildLinkedBackendId, options UnlinkBackendFromBuildOperationOptions) (*http.Request, error) {
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

// responderForUnlinkBackendFromBuild handles the response to the UnlinkBackendFromBuild request. The method always
// closes the http.Response Body.
func (c StaticSitesClient) responderForUnlinkBackendFromBuild(resp *http.Response) (result UnlinkBackendFromBuildOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

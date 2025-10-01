package bms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionContainersRefreshOperationResponse struct {
	HttpResponse *http.Response
}

type ProtectionContainersRefreshOperationOptions struct {
	Filter *string
}

func DefaultProtectionContainersRefreshOperationOptions() ProtectionContainersRefreshOperationOptions {
	return ProtectionContainersRefreshOperationOptions{}
}

func (o ProtectionContainersRefreshOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ProtectionContainersRefreshOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// ProtectionContainersRefresh ...
func (c BmsClient) ProtectionContainersRefresh(ctx context.Context, id BackupFabricId, options ProtectionContainersRefreshOperationOptions) (result ProtectionContainersRefreshOperationResponse, err error) {
	req, err := c.preparerForProtectionContainersRefresh(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ProtectionContainersRefresh", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ProtectionContainersRefresh", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProtectionContainersRefresh(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ProtectionContainersRefresh", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProtectionContainersRefresh prepares the ProtectionContainersRefresh request.
func (c BmsClient) preparerForProtectionContainersRefresh(ctx context.Context, id BackupFabricId, options ProtectionContainersRefreshOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/refreshContainers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForProtectionContainersRefresh handles the response to the ProtectionContainersRefresh request. The method always
// closes the http.Response Body.
func (c BmsClient) responderForProtectionContainersRefresh(resp *http.Response) (result ProtectionContainersRefreshOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

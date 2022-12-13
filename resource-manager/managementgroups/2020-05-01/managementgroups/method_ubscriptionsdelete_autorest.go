package managementgroups

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UbscriptionsDeleteOperationResponse struct {
	HttpResponse *http.Response
}

type UbscriptionsDeleteOperationOptions struct {
	CacheControl *string
}

func DefaultUbscriptionsDeleteOperationOptions() UbscriptionsDeleteOperationOptions {
	return UbscriptionsDeleteOperationOptions{}
}

func (o UbscriptionsDeleteOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.CacheControl != nil {
		out["Cache-Control"] = *o.CacheControl
	}

	return out
}

func (o UbscriptionsDeleteOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// UbscriptionsDelete ...
func (c ManagementGroupsClient) UbscriptionsDelete(ctx context.Context, id SubscriptionId, options UbscriptionsDeleteOperationOptions) (result UbscriptionsDeleteOperationResponse, err error) {
	req, err := c.preparerForUbscriptionsDelete(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "UbscriptionsDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "UbscriptionsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUbscriptionsDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "UbscriptionsDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUbscriptionsDelete prepares the UbscriptionsDelete request.
func (c ManagementGroupsClient) preparerForUbscriptionsDelete(ctx context.Context, id SubscriptionId, options UbscriptionsDeleteOperationOptions) (*http.Request, error) {
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

// responderForUbscriptionsDelete handles the response to the UbscriptionsDelete request. The method always
// closes the http.Response Body.
func (c ManagementGroupsClient) responderForUbscriptionsDelete(resp *http.Response) (result UbscriptionsDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

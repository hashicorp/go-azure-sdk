package managementgroups

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UbscriptionsCreateOperationResponse struct {
	HttpResponse *http.Response
	Model        *SubscriptionUnderManagementGroup
}

type UbscriptionsCreateOperationOptions struct {
	CacheControl *string
}

func DefaultUbscriptionsCreateOperationOptions() UbscriptionsCreateOperationOptions {
	return UbscriptionsCreateOperationOptions{}
}

func (o UbscriptionsCreateOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.CacheControl != nil {
		out["Cache-Control"] = *o.CacheControl
	}

	return out
}

func (o UbscriptionsCreateOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// UbscriptionsCreate ...
func (c ManagementGroupsClient) UbscriptionsCreate(ctx context.Context, id SubscriptionId, options UbscriptionsCreateOperationOptions) (result UbscriptionsCreateOperationResponse, err error) {
	req, err := c.preparerForUbscriptionsCreate(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "UbscriptionsCreate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "UbscriptionsCreate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUbscriptionsCreate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "UbscriptionsCreate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUbscriptionsCreate prepares the UbscriptionsCreate request.
func (c ManagementGroupsClient) preparerForUbscriptionsCreate(ctx context.Context, id SubscriptionId, options UbscriptionsCreateOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUbscriptionsCreate handles the response to the UbscriptionsCreate request. The method always
// closes the http.Response Body.
func (c ManagementGroupsClient) responderForUbscriptionsCreate(resp *http.Response) (result UbscriptionsCreateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

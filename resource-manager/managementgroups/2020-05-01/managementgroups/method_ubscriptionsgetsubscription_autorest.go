package managementgroups

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UbscriptionsGetSubscriptionOperationResponse struct {
	HttpResponse *http.Response
	Model        *SubscriptionUnderManagementGroup
}

type UbscriptionsGetSubscriptionOperationOptions struct {
	CacheControl *string
}

func DefaultUbscriptionsGetSubscriptionOperationOptions() UbscriptionsGetSubscriptionOperationOptions {
	return UbscriptionsGetSubscriptionOperationOptions{}
}

func (o UbscriptionsGetSubscriptionOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.CacheControl != nil {
		out["Cache-Control"] = *o.CacheControl
	}

	return out
}

func (o UbscriptionsGetSubscriptionOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// UbscriptionsGetSubscription ...
func (c ManagementGroupsClient) UbscriptionsGetSubscription(ctx context.Context, id SubscriptionId, options UbscriptionsGetSubscriptionOperationOptions) (result UbscriptionsGetSubscriptionOperationResponse, err error) {
	req, err := c.preparerForUbscriptionsGetSubscription(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "UbscriptionsGetSubscription", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "UbscriptionsGetSubscription", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUbscriptionsGetSubscription(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managementgroups.ManagementGroupsClient", "UbscriptionsGetSubscription", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUbscriptionsGetSubscription prepares the UbscriptionsGetSubscription request.
func (c ManagementGroupsClient) preparerForUbscriptionsGetSubscription(ctx context.Context, id SubscriptionId, options UbscriptionsGetSubscriptionOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUbscriptionsGetSubscription handles the response to the UbscriptionsGetSubscription request. The method always
// closes the http.Response Body.
func (c ManagementGroupsClient) responderForUbscriptionsGetSubscription(resp *http.Response) (result UbscriptionsGetSubscriptionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

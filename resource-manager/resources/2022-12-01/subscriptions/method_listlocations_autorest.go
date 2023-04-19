package subscriptions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListLocationsOperationResponse struct {
	HttpResponse *http.Response
	Model        *LocationListResult
}

type ListLocationsOperationOptions struct {
	IncludeExtendedLocations *bool
}

func DefaultListLocationsOperationOptions() ListLocationsOperationOptions {
	return ListLocationsOperationOptions{}
}

func (o ListLocationsOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListLocationsOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IncludeExtendedLocations != nil {
		out["includeExtendedLocations"] = *o.IncludeExtendedLocations
	}

	return out
}

// ListLocations ...
func (c SubscriptionsClient) ListLocations(ctx context.Context, id commonids.SubscriptionId, options ListLocationsOperationOptions) (result ListLocationsOperationResponse, err error) {
	req, err := c.preparerForListLocations(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "ListLocations", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "ListLocations", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListLocations(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "ListLocations", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListLocations prepares the ListLocations request.
func (c SubscriptionsClient) preparerForListLocations(ctx context.Context, id commonids.SubscriptionId, options ListLocationsOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/locations", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListLocations handles the response to the ListLocations request. The method always
// closes the http.Response Body.
func (c SubscriptionsClient) responderForListLocations(resp *http.Response) (result ListLocationsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

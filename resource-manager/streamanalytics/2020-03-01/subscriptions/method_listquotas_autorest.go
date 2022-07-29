package subscriptions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListQuotasOperationResponse struct {
	HttpResponse *http.Response
	Model        *SubscriptionQuotasListResult
}

// ListQuotas ...
func (c SubscriptionsClient) ListQuotas(ctx context.Context, id LocationId) (result ListQuotasOperationResponse, err error) {
	req, err := c.preparerForListQuotas(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "ListQuotas", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "ListQuotas", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListQuotas(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "ListQuotas", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListQuotas prepares the ListQuotas request.
func (c SubscriptionsClient) preparerForListQuotas(ctx context.Context, id LocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/quotas", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListQuotas handles the response to the ListQuotas request. The method always
// closes the http.Response Body.
func (c SubscriptionsClient) responderForListQuotas(resp *http.Response) (result ListQuotasOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

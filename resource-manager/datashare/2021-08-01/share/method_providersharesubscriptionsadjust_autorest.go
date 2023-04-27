package share

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderShareSubscriptionsAdjustOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProviderShareSubscription
}

// ProviderShareSubscriptionsAdjust ...
func (c ShareClient) ProviderShareSubscriptionsAdjust(ctx context.Context, id ProviderShareSubscriptionId, input ProviderShareSubscription) (result ProviderShareSubscriptionsAdjustOperationResponse, err error) {
	req, err := c.preparerForProviderShareSubscriptionsAdjust(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "share.ShareClient", "ProviderShareSubscriptionsAdjust", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "share.ShareClient", "ProviderShareSubscriptionsAdjust", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProviderShareSubscriptionsAdjust(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "share.ShareClient", "ProviderShareSubscriptionsAdjust", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProviderShareSubscriptionsAdjust prepares the ProviderShareSubscriptionsAdjust request.
func (c ShareClient) preparerForProviderShareSubscriptionsAdjust(ctx context.Context, id ProviderShareSubscriptionId, input ProviderShareSubscription) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/adjust", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForProviderShareSubscriptionsAdjust handles the response to the ProviderShareSubscriptionsAdjust request. The method always
// closes the http.Response Body.
func (c ShareClient) responderForProviderShareSubscriptionsAdjust(resp *http.Response) (result ProviderShareSubscriptionsAdjustOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

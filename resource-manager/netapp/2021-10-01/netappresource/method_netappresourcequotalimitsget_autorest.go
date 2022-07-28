package netappresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetAppResourceQuotaLimitsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *SubscriptionQuotaItem
}

// NetAppResourceQuotaLimitsGet ...
func (c NetAppResourceClient) NetAppResourceQuotaLimitsGet(ctx context.Context, id QuotaLimitId) (result NetAppResourceQuotaLimitsGetOperationResponse, err error) {
	req, err := c.preparerForNetAppResourceQuotaLimitsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceQuotaLimitsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceQuotaLimitsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNetAppResourceQuotaLimitsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceQuotaLimitsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNetAppResourceQuotaLimitsGet prepares the NetAppResourceQuotaLimitsGet request.
func (c NetAppResourceClient) preparerForNetAppResourceQuotaLimitsGet(ctx context.Context, id QuotaLimitId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForNetAppResourceQuotaLimitsGet handles the response to the NetAppResourceQuotaLimitsGet request. The method always
// closes the http.Response Body.
func (c NetAppResourceClient) responderForNetAppResourceQuotaLimitsGet(resp *http.Response) (result NetAppResourceQuotaLimitsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package netappresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetAppResourceQuotaLimitsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *SubscriptionQuotaItemList
}

// NetAppResourceQuotaLimitsList ...
func (c NetAppResourceClient) NetAppResourceQuotaLimitsList(ctx context.Context, id LocationId) (result NetAppResourceQuotaLimitsListOperationResponse, err error) {
	req, err := c.preparerForNetAppResourceQuotaLimitsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceQuotaLimitsList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceQuotaLimitsList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForNetAppResourceQuotaLimitsList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "netappresource.NetAppResourceClient", "NetAppResourceQuotaLimitsList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForNetAppResourceQuotaLimitsList prepares the NetAppResourceQuotaLimitsList request.
func (c NetAppResourceClient) preparerForNetAppResourceQuotaLimitsList(ctx context.Context, id LocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/quotaLimits", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForNetAppResourceQuotaLimitsList handles the response to the NetAppResourceQuotaLimitsList request. The method always
// closes the http.Response Body.
func (c NetAppResourceClient) responderForNetAppResourceQuotaLimitsList(resp *http.Response) (result NetAppResourceQuotaLimitsListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

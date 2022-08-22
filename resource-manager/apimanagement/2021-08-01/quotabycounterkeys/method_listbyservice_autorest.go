package quotabycounterkeys

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByServiceOperationResponse struct {
	HttpResponse *http.Response
	Model        *QuotaCounterCollection
}

// ListByService ...
func (c QuotaByCounterKeysClient) ListByService(ctx context.Context, id QuotaId) (result ListByServiceOperationResponse, err error) {
	req, err := c.preparerForListByService(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "quotabycounterkeys.QuotaByCounterKeysClient", "ListByService", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "quotabycounterkeys.QuotaByCounterKeysClient", "ListByService", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByService(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "quotabycounterkeys.QuotaByCounterKeysClient", "ListByService", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByService prepares the ListByService request.
func (c QuotaByCounterKeysClient) preparerForListByService(ctx context.Context, id QuotaId) (*http.Request, error) {
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

// responderForListByService handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (c QuotaByCounterKeysClient) responderForListByService(resp *http.Response) (result ListByServiceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

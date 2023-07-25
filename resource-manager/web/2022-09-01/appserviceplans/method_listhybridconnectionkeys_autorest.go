package appserviceplans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListHybridConnectionKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *HybridConnectionKey
}

// ListHybridConnectionKeys ...
func (c AppServicePlansClient) ListHybridConnectionKeys(ctx context.Context, id HybridConnectionNamespaceRelayId) (result ListHybridConnectionKeysOperationResponse, err error) {
	req, err := c.preparerForListHybridConnectionKeys(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListHybridConnectionKeys", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListHybridConnectionKeys", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListHybridConnectionKeys(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceplans.AppServicePlansClient", "ListHybridConnectionKeys", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListHybridConnectionKeys prepares the ListHybridConnectionKeys request.
func (c AppServicePlansClient) preparerForListHybridConnectionKeys(ctx context.Context, id HybridConnectionNamespaceRelayId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listKeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListHybridConnectionKeys handles the response to the ListHybridConnectionKeys request. The method always
// closes the http.Response Body.
func (c AppServicePlansClient) responderForListHybridConnectionKeys(resp *http.Response) (result ListHybridConnectionKeysOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package kusto

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkusListOperationResponse struct {
	HttpResponse *http.Response
	Model        *SkuDescriptionList
}

// SkusList ...
func (c KustoClient) SkusList(ctx context.Context, id LocationId) (result SkusListOperationResponse, err error) {
	req, err := c.preparerForSkusList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.KustoClient", "SkusList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.KustoClient", "SkusList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSkusList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.KustoClient", "SkusList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSkusList prepares the SkusList request.
func (c KustoClient) preparerForSkusList(ctx context.Context, id LocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/skus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSkusList handles the response to the SkusList request. The method always
// closes the http.Response Body.
func (c KustoClient) responderForSkusList(resp *http.Response) (result SkusListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

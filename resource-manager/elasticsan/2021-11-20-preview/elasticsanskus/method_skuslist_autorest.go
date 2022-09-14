package elasticsanskus

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

type SkusListOperationResponse struct {
	HttpResponse *http.Response
	Model        *SkuInformationList
}

type SkusListOperationOptions struct {
	Filter *string
}

func DefaultSkusListOperationOptions() SkusListOperationOptions {
	return SkusListOperationOptions{}
}

func (o SkusListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o SkusListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// SkusList ...
func (c ElasticSanSkusClient) SkusList(ctx context.Context, id commonids.SubscriptionId, options SkusListOperationOptions) (result SkusListOperationResponse, err error) {
	req, err := c.preparerForSkusList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elasticsanskus.ElasticSanSkusClient", "SkusList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "elasticsanskus.ElasticSanSkusClient", "SkusList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSkusList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "elasticsanskus.ElasticSanSkusClient", "SkusList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSkusList prepares the SkusList request.
func (c ElasticSanSkusClient) preparerForSkusList(ctx context.Context, id commonids.SubscriptionId, options SkusListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.ElasticSan/skus", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSkusList handles the response to the SkusList request. The method always
// closes the http.Response Body.
func (c ElasticSanSkusClient) responderForSkusList(resp *http.Response) (result SkusListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

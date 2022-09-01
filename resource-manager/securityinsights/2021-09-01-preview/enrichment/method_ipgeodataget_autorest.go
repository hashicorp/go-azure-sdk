package enrichment

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

type IPGeodataGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *EnrichmentIPGeodata
}

type IPGeodataGetOperationOptions struct {
	IPAddress *string
}

func DefaultIPGeodataGetOperationOptions() IPGeodataGetOperationOptions {
	return IPGeodataGetOperationOptions{}
}

func (o IPGeodataGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o IPGeodataGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IPAddress != nil {
		out["ipAddress"] = *o.IPAddress
	}

	return out
}

// IPGeodataGet ...
func (c EnrichmentClient) IPGeodataGet(ctx context.Context, id commonids.ResourceGroupId, options IPGeodataGetOperationOptions) (result IPGeodataGetOperationResponse, err error) {
	req, err := c.preparerForIPGeodataGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enrichment.EnrichmentClient", "IPGeodataGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "enrichment.EnrichmentClient", "IPGeodataGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIPGeodataGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enrichment.EnrichmentClient", "IPGeodataGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIPGeodataGet prepares the IPGeodataGet request.
func (c EnrichmentClient) preparerForIPGeodataGet(ctx context.Context, id commonids.ResourceGroupId, options IPGeodataGetOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/enrichment/ip/geodata", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIPGeodataGet handles the response to the IPGeodataGet request. The method always
// closes the http.Response Body.
func (c EnrichmentClient) responderForIPGeodataGet(resp *http.Response) (result IPGeodataGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

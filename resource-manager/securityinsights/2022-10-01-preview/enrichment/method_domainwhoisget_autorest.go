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

type DomainWhoisGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *EnrichmentDomainWhois
}

type DomainWhoisGetOperationOptions struct {
	Domain *string
}

func DefaultDomainWhoisGetOperationOptions() DomainWhoisGetOperationOptions {
	return DomainWhoisGetOperationOptions{}
}

func (o DomainWhoisGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o DomainWhoisGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Domain != nil {
		out["domain"] = *o.Domain
	}

	return out
}

// DomainWhoisGet ...
func (c EnrichmentClient) DomainWhoisGet(ctx context.Context, id commonids.ResourceGroupId, options DomainWhoisGetOperationOptions) (result DomainWhoisGetOperationResponse, err error) {
	req, err := c.preparerForDomainWhoisGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enrichment.EnrichmentClient", "DomainWhoisGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "enrichment.EnrichmentClient", "DomainWhoisGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDomainWhoisGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "enrichment.EnrichmentClient", "DomainWhoisGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDomainWhoisGet prepares the DomainWhoisGet request.
func (c EnrichmentClient) preparerForDomainWhoisGet(ctx context.Context, id commonids.ResourceGroupId, options DomainWhoisGetOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/enrichment/domain/whois", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDomainWhoisGet handles the response to the DomainWhoisGet request. The method always
// closes the http.Response Body.
func (c EnrichmentClient) responderForDomainWhoisGet(resp *http.Response) (result DomainWhoisGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

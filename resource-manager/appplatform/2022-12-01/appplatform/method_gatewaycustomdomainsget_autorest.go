package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayCustomDomainsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *GatewayCustomDomainResource
}

// GatewayCustomDomainsGet ...
func (c AppPlatformClient) GatewayCustomDomainsGet(ctx context.Context, id GatewayDomainId) (result GatewayCustomDomainsGetOperationResponse, err error) {
	req, err := c.preparerForGatewayCustomDomainsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayCustomDomainsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayCustomDomainsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGatewayCustomDomainsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayCustomDomainsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGatewayCustomDomainsGet prepares the GatewayCustomDomainsGet request.
func (c AppPlatformClient) preparerForGatewayCustomDomainsGet(ctx context.Context, id GatewayDomainId) (*http.Request, error) {
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

// responderForGatewayCustomDomainsGet handles the response to the GatewayCustomDomainsGet request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForGatewayCustomDomainsGet(resp *http.Response) (result GatewayCustomDomainsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

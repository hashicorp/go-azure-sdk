package appplatform

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiPortalCustomDomainsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ApiPortalCustomDomainResource
}

// ApiPortalCustomDomainsGet ...
func (c AppPlatformClient) ApiPortalCustomDomainsGet(ctx context.Context, id ApiPortalDomainId) (result ApiPortalCustomDomainsGetOperationResponse, err error) {
	req, err := c.preparerForApiPortalCustomDomainsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalCustomDomainsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalCustomDomainsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForApiPortalCustomDomainsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalCustomDomainsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForApiPortalCustomDomainsGet prepares the ApiPortalCustomDomainsGet request.
func (c AppPlatformClient) preparerForApiPortalCustomDomainsGet(ctx context.Context, id ApiPortalDomainId) (*http.Request, error) {
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

// responderForApiPortalCustomDomainsGet handles the response to the ApiPortalCustomDomainsGet request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForApiPortalCustomDomainsGet(resp *http.Response) (result ApiPortalCustomDomainsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package marketplaceregistrationdefinitions

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WithoutScopeGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *MarketplaceRegistrationDefinition
}

// WithoutScopeGet ...
func (c MarketplaceRegistrationDefinitionsClient) WithoutScopeGet(ctx context.Context, id MarketplaceRegistrationDefinitionId) (result WithoutScopeGetOperationResponse, err error) {
	req, err := c.preparerForWithoutScopeGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceregistrationdefinitions.MarketplaceRegistrationDefinitionsClient", "WithoutScopeGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceregistrationdefinitions.MarketplaceRegistrationDefinitionsClient", "WithoutScopeGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWithoutScopeGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "marketplaceregistrationdefinitions.MarketplaceRegistrationDefinitionsClient", "WithoutScopeGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWithoutScopeGet prepares the WithoutScopeGet request.
func (c MarketplaceRegistrationDefinitionsClient) preparerForWithoutScopeGet(ctx context.Context, id MarketplaceRegistrationDefinitionId) (*http.Request, error) {
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

// responderForWithoutScopeGet handles the response to the WithoutScopeGet request. The method always
// closes the http.Response Body.
func (c MarketplaceRegistrationDefinitionsClient) responderForWithoutScopeGet(resp *http.Response) (result WithoutScopeGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package vaultusages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsagesListByVaultsOperationResponse struct {
	HttpResponse *http.Response
	Model        *VaultUsageList
}

// UsagesListByVaults ...
func (c VaultUsagesClient) UsagesListByVaults(ctx context.Context, id VaultId) (result UsagesListByVaultsOperationResponse, err error) {
	req, err := c.preparerForUsagesListByVaults(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vaultusages.VaultUsagesClient", "UsagesListByVaults", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "vaultusages.VaultUsagesClient", "UsagesListByVaults", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUsagesListByVaults(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vaultusages.VaultUsagesClient", "UsagesListByVaults", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUsagesListByVaults prepares the UsagesListByVaults request.
func (c VaultUsagesClient) preparerForUsagesListByVaults(ctx context.Context, id VaultId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/usages", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUsagesListByVaults handles the response to the UsagesListByVaults request. The method always
// closes the http.Response Body.
func (c VaultUsagesClient) responderForUsagesListByVaults(resp *http.Response) (result UsagesListByVaultsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

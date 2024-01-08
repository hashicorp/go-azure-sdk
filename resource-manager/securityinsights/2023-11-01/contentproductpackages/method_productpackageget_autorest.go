package contentproductpackages

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductPackageGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProductPackageModel
}

// ProductPackageGet ...
func (c ContentProductPackagesClient) ProductPackageGet(ctx context.Context, id ContentProductPackageId) (result ProductPackageGetOperationResponse, err error) {
	req, err := c.preparerForProductPackageGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentproductpackages.ContentProductPackagesClient", "ProductPackageGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentproductpackages.ContentProductPackagesClient", "ProductPackageGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProductPackageGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentproductpackages.ContentProductPackagesClient", "ProductPackageGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProductPackageGet prepares the ProductPackageGet request.
func (c ContentProductPackagesClient) preparerForProductPackageGet(ctx context.Context, id ContentProductPackageId) (*http.Request, error) {
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

// responderForProductPackageGet handles the response to the ProductPackageGet request. The method always
// closes the http.Response Body.
func (c ContentProductPackagesClient) responderForProductPackageGet(resp *http.Response) (result ProductPackageGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

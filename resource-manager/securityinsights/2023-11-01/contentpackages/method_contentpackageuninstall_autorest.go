package contentpackages

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentPackageUninstallOperationResponse struct {
	HttpResponse *http.Response
}

// ContentPackageUninstall ...
func (c ContentPackagesClient) ContentPackageUninstall(ctx context.Context, id ContentPackageId) (result ContentPackageUninstallOperationResponse, err error) {
	req, err := c.preparerForContentPackageUninstall(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentpackages.ContentPackagesClient", "ContentPackageUninstall", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentpackages.ContentPackagesClient", "ContentPackageUninstall", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContentPackageUninstall(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentpackages.ContentPackagesClient", "ContentPackageUninstall", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContentPackageUninstall prepares the ContentPackageUninstall request.
func (c ContentPackagesClient) preparerForContentPackageUninstall(ctx context.Context, id ContentPackageId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForContentPackageUninstall handles the response to the ContentPackageUninstall request. The method always
// closes the http.Response Body.
func (c ContentPackagesClient) responderForContentPackageUninstall(resp *http.Response) (result ContentPackageUninstallOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

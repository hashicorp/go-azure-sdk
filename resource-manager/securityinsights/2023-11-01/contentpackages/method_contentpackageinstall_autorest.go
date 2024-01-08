package contentpackages

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentPackageInstallOperationResponse struct {
	HttpResponse *http.Response
	Model        *PackageModel
}

// ContentPackageInstall ...
func (c ContentPackagesClient) ContentPackageInstall(ctx context.Context, id ContentPackageId, input PackageModel) (result ContentPackageInstallOperationResponse, err error) {
	req, err := c.preparerForContentPackageInstall(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentpackages.ContentPackagesClient", "ContentPackageInstall", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentpackages.ContentPackagesClient", "ContentPackageInstall", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForContentPackageInstall(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "contentpackages.ContentPackagesClient", "ContentPackageInstall", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForContentPackageInstall prepares the ContentPackageInstall request.
func (c ContentPackagesClient) preparerForContentPackageInstall(ctx context.Context, id ContentPackageId, input PackageModel) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForContentPackageInstall handles the response to the ContentPackageInstall request. The method always
// closes the http.Response Body.
func (c ContentPackagesClient) responderForContentPackageInstall(resp *http.Response) (result ContentPackageInstallOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

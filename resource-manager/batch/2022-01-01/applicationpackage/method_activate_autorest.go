package applicationpackage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ApplicationPackage
}

// Activate ...
func (c ApplicationPackageClient) Activate(ctx context.Context, id VersionId, input ActivateApplicationPackageParameters) (result ActivateOperationResponse, err error) {
	req, err := c.preparerForActivate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationpackage.ApplicationPackageClient", "Activate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationpackage.ApplicationPackageClient", "Activate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForActivate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationpackage.ApplicationPackageClient", "Activate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForActivate prepares the Activate request.
func (c ApplicationPackageClient) preparerForActivate(ctx context.Context, id VersionId, input ActivateApplicationPackageParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/activate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForActivate handles the response to the Activate request. The method always
// closes the http.Response Body.
func (c ApplicationPackageClient) responderForActivate(resp *http.Response) (result ActivateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

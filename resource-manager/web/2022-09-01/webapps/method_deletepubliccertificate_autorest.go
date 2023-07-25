package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePublicCertificateOperationResponse struct {
	HttpResponse *http.Response
}

// DeletePublicCertificate ...
func (c WebAppsClient) DeletePublicCertificate(ctx context.Context, id PublicCertificateId) (result DeletePublicCertificateOperationResponse, err error) {
	req, err := c.preparerForDeletePublicCertificate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeletePublicCertificate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeletePublicCertificate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeletePublicCertificate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeletePublicCertificate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeletePublicCertificate prepares the DeletePublicCertificate request.
func (c WebAppsClient) preparerForDeletePublicCertificate(ctx context.Context, id PublicCertificateId) (*http.Request, error) {
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

// responderForDeletePublicCertificate handles the response to the DeletePublicCertificate request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeletePublicCertificate(resp *http.Response) (result DeletePublicCertificateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

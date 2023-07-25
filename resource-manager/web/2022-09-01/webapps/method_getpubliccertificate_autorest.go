package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPublicCertificateOperationResponse struct {
	HttpResponse *http.Response
	Model        *PublicCertificate
}

// GetPublicCertificate ...
func (c WebAppsClient) GetPublicCertificate(ctx context.Context, id PublicCertificateId) (result GetPublicCertificateOperationResponse, err error) {
	req, err := c.preparerForGetPublicCertificate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPublicCertificate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPublicCertificate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetPublicCertificate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetPublicCertificate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetPublicCertificate prepares the GetPublicCertificate request.
func (c WebAppsClient) preparerForGetPublicCertificate(ctx context.Context, id PublicCertificateId) (*http.Request, error) {
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

// responderForGetPublicCertificate handles the response to the GetPublicCertificate request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetPublicCertificate(resp *http.Response) (result GetPublicCertificateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

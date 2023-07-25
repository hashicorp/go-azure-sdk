package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdatePublicCertificateOperationResponse struct {
	HttpResponse *http.Response
	Model        *PublicCertificate
}

// CreateOrUpdatePublicCertificate ...
func (c WebAppsClient) CreateOrUpdatePublicCertificate(ctx context.Context, id PublicCertificateId, input PublicCertificate) (result CreateOrUpdatePublicCertificateOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdatePublicCertificate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdatePublicCertificate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdatePublicCertificate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdatePublicCertificate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdatePublicCertificate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdatePublicCertificate prepares the CreateOrUpdatePublicCertificate request.
func (c WebAppsClient) preparerForCreateOrUpdatePublicCertificate(ctx context.Context, id PublicCertificateId, input PublicCertificate) (*http.Request, error) {
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

// responderForCreateOrUpdatePublicCertificate handles the response to the CreateOrUpdatePublicCertificate request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdatePublicCertificate(resp *http.Response) (result CreateOrUpdatePublicCertificateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

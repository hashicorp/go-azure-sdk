package appservicecertificateorders

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteCertificateOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteCertificate ...
func (c AppServiceCertificateOrdersClient) DeleteCertificate(ctx context.Context, id CertificateId) (result DeleteCertificateOperationResponse, err error) {
	req, err := c.preparerForDeleteCertificate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "DeleteCertificate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "DeleteCertificate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteCertificate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "DeleteCertificate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteCertificate prepares the DeleteCertificate request.
func (c AppServiceCertificateOrdersClient) preparerForDeleteCertificate(ctx context.Context, id CertificateId) (*http.Request, error) {
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

// responderForDeleteCertificate handles the response to the DeleteCertificate request. The method always
// closes the http.Response Body.
func (c AppServiceCertificateOrdersClient) responderForDeleteCertificate(resp *http.Response) (result DeleteCertificateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

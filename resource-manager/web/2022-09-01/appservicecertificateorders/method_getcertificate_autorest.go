package appservicecertificateorders

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCertificateOperationResponse struct {
	HttpResponse *http.Response
	Model        *AppServiceCertificateResource
}

// GetCertificate ...
func (c AppServiceCertificateOrdersClient) GetCertificate(ctx context.Context, id CertificateId) (result GetCertificateOperationResponse, err error) {
	req, err := c.preparerForGetCertificate(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "GetCertificate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "GetCertificate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetCertificate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "GetCertificate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetCertificate prepares the GetCertificate request.
func (c AppServiceCertificateOrdersClient) preparerForGetCertificate(ctx context.Context, id CertificateId) (*http.Request, error) {
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

// responderForGetCertificate handles the response to the GetCertificate request. The method always
// closes the http.Response Body.
func (c AppServiceCertificateOrdersClient) responderForGetCertificate(resp *http.Response) (result GetCertificateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package appservicecertificateorders

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCertificateOperationResponse struct {
	HttpResponse *http.Response
	Model        *AppServiceCertificateResource
}

// UpdateCertificate ...
func (c AppServiceCertificateOrdersClient) UpdateCertificate(ctx context.Context, id CertificateId, input AppServiceCertificatePatchResource) (result UpdateCertificateOperationResponse, err error) {
	req, err := c.preparerForUpdateCertificate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "UpdateCertificate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "UpdateCertificate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateCertificate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appservicecertificateorders.AppServiceCertificateOrdersClient", "UpdateCertificate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateCertificate prepares the UpdateCertificate request.
func (c AppServiceCertificateOrdersClient) preparerForUpdateCertificate(ctx context.Context, id CertificateId, input AppServiceCertificatePatchResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateCertificate handles the response to the UpdateCertificate request. The method always
// closes the http.Response Body.
func (c AppServiceCertificateOrdersClient) responderForUpdateCertificate(resp *http.Response) (result UpdateCertificateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package connectedenvironments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificatesDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// CertificatesDelete ...
func (c ConnectedEnvironmentsClient) CertificatesDelete(ctx context.Context, id ConnectedEnvironmentCertificateId) (result CertificatesDeleteOperationResponse, err error) {
	req, err := c.preparerForCertificatesDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "connectedenvironments.ConnectedEnvironmentsClient", "CertificatesDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "connectedenvironments.ConnectedEnvironmentsClient", "CertificatesDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCertificatesDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "connectedenvironments.ConnectedEnvironmentsClient", "CertificatesDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCertificatesDelete prepares the CertificatesDelete request.
func (c ConnectedEnvironmentsClient) preparerForCertificatesDelete(ctx context.Context, id ConnectedEnvironmentCertificateId) (*http.Request, error) {
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

// responderForCertificatesDelete handles the response to the CertificatesDelete request. The method always
// closes the http.Response Body.
func (c ConnectedEnvironmentsClient) responderForCertificatesDelete(resp *http.Response) (result CertificatesDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

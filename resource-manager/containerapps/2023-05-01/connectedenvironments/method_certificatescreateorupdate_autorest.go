package connectedenvironments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificatesCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *Certificate
}

// CertificatesCreateOrUpdate ...
func (c ConnectedEnvironmentsClient) CertificatesCreateOrUpdate(ctx context.Context, id ConnectedEnvironmentCertificateId, input Certificate) (result CertificatesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForCertificatesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "connectedenvironments.ConnectedEnvironmentsClient", "CertificatesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "connectedenvironments.ConnectedEnvironmentsClient", "CertificatesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCertificatesCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "connectedenvironments.ConnectedEnvironmentsClient", "CertificatesCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCertificatesCreateOrUpdate prepares the CertificatesCreateOrUpdate request.
func (c ConnectedEnvironmentsClient) preparerForCertificatesCreateOrUpdate(ctx context.Context, id ConnectedEnvironmentCertificateId, input Certificate) (*http.Request, error) {
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

// responderForCertificatesCreateOrUpdate handles the response to the CertificatesCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c ConnectedEnvironmentsClient) responderForCertificatesCreateOrUpdate(resp *http.Response) (result CertificatesCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

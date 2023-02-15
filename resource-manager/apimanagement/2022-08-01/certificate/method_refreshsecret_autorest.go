package certificate

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RefreshSecretOperationResponse struct {
	HttpResponse *http.Response
	Model        *CertificateContract
}

// RefreshSecret ...
func (c CertificateClient) RefreshSecret(ctx context.Context, id CertificateId) (result RefreshSecretOperationResponse, err error) {
	req, err := c.preparerForRefreshSecret(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "certificate.CertificateClient", "RefreshSecret", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "certificate.CertificateClient", "RefreshSecret", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRefreshSecret(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "certificate.CertificateClient", "RefreshSecret", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRefreshSecret prepares the RefreshSecret request.
func (c CertificateClient) preparerForRefreshSecret(ctx context.Context, id CertificateId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/refreshSecret", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRefreshSecret handles the response to the RefreshSecret request. The method always
// closes the http.Response Body.
func (c CertificateClient) responderForRefreshSecret(resp *http.Response) (result RefreshSecretOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

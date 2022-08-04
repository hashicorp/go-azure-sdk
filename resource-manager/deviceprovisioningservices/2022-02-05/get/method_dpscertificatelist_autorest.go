package get

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DpsCertificateListOperationResponse struct {
	HttpResponse *http.Response
	Model        *CertificateListDescription
}

// DpsCertificateList ...
func (c GETClient) DpsCertificateList(ctx context.Context, id ProvisioningServiceId) (result DpsCertificateListOperationResponse, err error) {
	req, err := c.preparerForDpsCertificateList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "DpsCertificateList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "DpsCertificateList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDpsCertificateList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "get.GETClient", "DpsCertificateList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDpsCertificateList prepares the DpsCertificateList request.
func (c GETClient) preparerForDpsCertificateList(ctx context.Context, id ProvisioningServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/certificates", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDpsCertificateList handles the response to the DpsCertificateList request. The method always
// closes the http.Response Body.
func (c GETClient) responderForDpsCertificateList(resp *http.Response) (result DpsCertificateListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

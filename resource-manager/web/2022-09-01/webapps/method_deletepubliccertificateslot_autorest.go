package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePublicCertificateSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeletePublicCertificateSlot ...
func (c WebAppsClient) DeletePublicCertificateSlot(ctx context.Context, id SlotPublicCertificateId) (result DeletePublicCertificateSlotOperationResponse, err error) {
	req, err := c.preparerForDeletePublicCertificateSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeletePublicCertificateSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeletePublicCertificateSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeletePublicCertificateSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeletePublicCertificateSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeletePublicCertificateSlot prepares the DeletePublicCertificateSlot request.
func (c WebAppsClient) preparerForDeletePublicCertificateSlot(ctx context.Context, id SlotPublicCertificateId) (*http.Request, error) {
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

// responderForDeletePublicCertificateSlot handles the response to the DeletePublicCertificateSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeletePublicCertificateSlot(resp *http.Response) (result DeletePublicCertificateSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdatePublicCertificateSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *PublicCertificate
}

// CreateOrUpdatePublicCertificateSlot ...
func (c WebAppsClient) CreateOrUpdatePublicCertificateSlot(ctx context.Context, id SlotPublicCertificateId, input PublicCertificate) (result CreateOrUpdatePublicCertificateSlotOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdatePublicCertificateSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdatePublicCertificateSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdatePublicCertificateSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdatePublicCertificateSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "CreateOrUpdatePublicCertificateSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdatePublicCertificateSlot prepares the CreateOrUpdatePublicCertificateSlot request.
func (c WebAppsClient) preparerForCreateOrUpdatePublicCertificateSlot(ctx context.Context, id SlotPublicCertificateId, input PublicCertificate) (*http.Request, error) {
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

// responderForCreateOrUpdatePublicCertificateSlot handles the response to the CreateOrUpdatePublicCertificateSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForCreateOrUpdatePublicCertificateSlot(resp *http.Response) (result CreateOrUpdatePublicCertificateSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

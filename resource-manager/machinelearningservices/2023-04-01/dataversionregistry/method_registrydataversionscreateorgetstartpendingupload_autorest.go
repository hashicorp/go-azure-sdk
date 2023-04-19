package dataversionregistry

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryDataVersionsCreateOrGetStartPendingUploadOperationResponse struct {
	HttpResponse *http.Response
	Model        *PendingUploadResponseDto
}

// RegistryDataVersionsCreateOrGetStartPendingUpload ...
func (c DataVersionRegistryClient) RegistryDataVersionsCreateOrGetStartPendingUpload(ctx context.Context, id VersionId, input PendingUploadRequestDto) (result RegistryDataVersionsCreateOrGetStartPendingUploadOperationResponse, err error) {
	req, err := c.preparerForRegistryDataVersionsCreateOrGetStartPendingUpload(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsCreateOrGetStartPendingUpload", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsCreateOrGetStartPendingUpload", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistryDataVersionsCreateOrGetStartPendingUpload(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsCreateOrGetStartPendingUpload", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistryDataVersionsCreateOrGetStartPendingUpload prepares the RegistryDataVersionsCreateOrGetStartPendingUpload request.
func (c DataVersionRegistryClient) preparerForRegistryDataVersionsCreateOrGetStartPendingUpload(ctx context.Context, id VersionId, input PendingUploadRequestDto) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/startPendingUpload", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRegistryDataVersionsCreateOrGetStartPendingUpload handles the response to the RegistryDataVersionsCreateOrGetStartPendingUpload request. The method always
// closes the http.Response Body.
func (c DataVersionRegistryClient) responderForRegistryDataVersionsCreateOrGetStartPendingUpload(resp *http.Response) (result RegistryDataVersionsCreateOrGetStartPendingUploadOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

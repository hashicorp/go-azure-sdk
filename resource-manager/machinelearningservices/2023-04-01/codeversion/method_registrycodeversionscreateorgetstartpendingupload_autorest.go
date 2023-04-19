package codeversion

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryCodeVersionsCreateOrGetStartPendingUploadOperationResponse struct {
	HttpResponse *http.Response
	Model        *PendingUploadResponseDto
}

// RegistryCodeVersionsCreateOrGetStartPendingUpload ...
func (c CodeVersionClient) RegistryCodeVersionsCreateOrGetStartPendingUpload(ctx context.Context, id RegistryCodeVersionId, input PendingUploadRequestDto) (result RegistryCodeVersionsCreateOrGetStartPendingUploadOperationResponse, err error) {
	req, err := c.preparerForRegistryCodeVersionsCreateOrGetStartPendingUpload(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsCreateOrGetStartPendingUpload", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsCreateOrGetStartPendingUpload", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistryCodeVersionsCreateOrGetStartPendingUpload(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsCreateOrGetStartPendingUpload", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistryCodeVersionsCreateOrGetStartPendingUpload prepares the RegistryCodeVersionsCreateOrGetStartPendingUpload request.
func (c CodeVersionClient) preparerForRegistryCodeVersionsCreateOrGetStartPendingUpload(ctx context.Context, id RegistryCodeVersionId, input PendingUploadRequestDto) (*http.Request, error) {
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

// responderForRegistryCodeVersionsCreateOrGetStartPendingUpload handles the response to the RegistryCodeVersionsCreateOrGetStartPendingUpload request. The method always
// closes the http.Response Body.
func (c CodeVersionClient) responderForRegistryCodeVersionsCreateOrGetStartPendingUpload(resp *http.Response) (result RegistryCodeVersionsCreateOrGetStartPendingUploadOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package modelversion

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryModelVersionsCreateOrGetStartPendingUploadOperationResponse struct {
	HttpResponse *http.Response
	Model        *PendingUploadResponseDto
}

// RegistryModelVersionsCreateOrGetStartPendingUpload ...
func (c ModelVersionClient) RegistryModelVersionsCreateOrGetStartPendingUpload(ctx context.Context, id RegistryModelVersionId, input PendingUploadRequestDto) (result RegistryModelVersionsCreateOrGetStartPendingUploadOperationResponse, err error) {
	req, err := c.preparerForRegistryModelVersionsCreateOrGetStartPendingUpload(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsCreateOrGetStartPendingUpload", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsCreateOrGetStartPendingUpload", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegistryModelVersionsCreateOrGetStartPendingUpload(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsCreateOrGetStartPendingUpload", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegistryModelVersionsCreateOrGetStartPendingUpload prepares the RegistryModelVersionsCreateOrGetStartPendingUpload request.
func (c ModelVersionClient) preparerForRegistryModelVersionsCreateOrGetStartPendingUpload(ctx context.Context, id RegistryModelVersionId, input PendingUploadRequestDto) (*http.Request, error) {
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

// responderForRegistryModelVersionsCreateOrGetStartPendingUpload handles the response to the RegistryModelVersionsCreateOrGetStartPendingUpload request. The method always
// closes the http.Response Body.
func (c ModelVersionClient) responderForRegistryModelVersionsCreateOrGetStartPendingUpload(resp *http.Response) (result RegistryModelVersionsCreateOrGetStartPendingUploadOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

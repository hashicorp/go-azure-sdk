package iotsecuritysolutions

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotSecuritySolutionGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *IoTSecuritySolutionModel
}

// IotSecuritySolutionGet ...
func (c IotSecuritySolutionsClient) IotSecuritySolutionGet(ctx context.Context, id IotSecuritySolutionId) (result IotSecuritySolutionGetOperationResponse, err error) {
	req, err := c.preparerForIotSecuritySolutionGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IotSecuritySolutionGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IotSecuritySolutionGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIotSecuritySolutionGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IotSecuritySolutionGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIotSecuritySolutionGet prepares the IotSecuritySolutionGet request.
func (c IotSecuritySolutionsClient) preparerForIotSecuritySolutionGet(ctx context.Context, id IotSecuritySolutionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIotSecuritySolutionGet handles the response to the IotSecuritySolutionGet request. The method always
// closes the http.Response Body.
func (c IotSecuritySolutionsClient) responderForIotSecuritySolutionGet(resp *http.Response) (result IotSecuritySolutionGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

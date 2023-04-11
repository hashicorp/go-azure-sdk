package iotsecuritysolutions

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotSecuritySolutionCreateOperationResponse struct {
	HttpResponse *http.Response
	Model        *IoTSecuritySolutionModel
}

// IotSecuritySolutionCreate ...
func (c IotSecuritySolutionsClient) IotSecuritySolutionCreate(ctx context.Context, id IotSecuritySolutionId, input IoTSecuritySolutionModel) (result IotSecuritySolutionCreateOperationResponse, err error) {
	req, err := c.preparerForIotSecuritySolutionCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IotSecuritySolutionCreate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IotSecuritySolutionCreate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIotSecuritySolutionCreate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "iotsecuritysolutions.IotSecuritySolutionsClient", "IotSecuritySolutionCreate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIotSecuritySolutionCreate prepares the IotSecuritySolutionCreate request.
func (c IotSecuritySolutionsClient) preparerForIotSecuritySolutionCreate(ctx context.Context, id IotSecuritySolutionId, input IoTSecuritySolutionModel) (*http.Request, error) {
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

// responderForIotSecuritySolutionCreate handles the response to the IotSecuritySolutionCreate request. The method always
// closes the http.Response Body.
func (c IotSecuritySolutionsClient) responderForIotSecuritySolutionCreate(resp *http.Response) (result IotSecuritySolutionCreateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

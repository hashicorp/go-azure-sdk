package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeploymentSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *Deployment
}

// GetDeploymentSlot ...
func (c WebAppsClient) GetDeploymentSlot(ctx context.Context, id SlotDeploymentId) (result GetDeploymentSlotOperationResponse, err error) {
	req, err := c.preparerForGetDeploymentSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDeploymentSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDeploymentSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDeploymentSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetDeploymentSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDeploymentSlot prepares the GetDeploymentSlot request.
func (c WebAppsClient) preparerForGetDeploymentSlot(ctx context.Context, id SlotDeploymentId) (*http.Request, error) {
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

// responderForGetDeploymentSlot handles the response to the GetDeploymentSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetDeploymentSlot(resp *http.Response) (result GetDeploymentSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

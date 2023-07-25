package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteDeploymentSlotOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteDeploymentSlot ...
func (c WebAppsClient) DeleteDeploymentSlot(ctx context.Context, id SlotDeploymentId) (result DeleteDeploymentSlotOperationResponse, err error) {
	req, err := c.preparerForDeleteDeploymentSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteDeploymentSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteDeploymentSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteDeploymentSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "DeleteDeploymentSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteDeploymentSlot prepares the DeleteDeploymentSlot request.
func (c WebAppsClient) preparerForDeleteDeploymentSlot(ctx context.Context, id SlotDeploymentId) (*http.Request, error) {
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

// responderForDeleteDeploymentSlot handles the response to the DeleteDeploymentSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForDeleteDeploymentSlot(resp *http.Response) (result DeleteDeploymentSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

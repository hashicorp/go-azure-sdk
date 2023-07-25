package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDeploymentLogSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *Deployment
}

// ListDeploymentLogSlot ...
func (c WebAppsClient) ListDeploymentLogSlot(ctx context.Context, id SlotDeploymentId) (result ListDeploymentLogSlotOperationResponse, err error) {
	req, err := c.preparerForListDeploymentLogSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeploymentLogSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeploymentLogSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListDeploymentLogSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListDeploymentLogSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListDeploymentLogSlot prepares the ListDeploymentLogSlot request.
func (c WebAppsClient) preparerForListDeploymentLogSlot(ctx context.Context, id SlotDeploymentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/log", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListDeploymentLogSlot handles the response to the ListDeploymentLogSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForListDeploymentLogSlot(resp *http.Response) (result ListDeploymentLogSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

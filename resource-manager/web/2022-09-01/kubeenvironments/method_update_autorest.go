package kubeenvironments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *KubeEnvironment
}

// Update ...
func (c KubeEnvironmentsClient) Update(ctx context.Context, id KubeEnvironmentId, input KubeEnvironmentPatchResource) (result UpdateOperationResponse, err error) {
	req, err := c.preparerForUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubeenvironments.KubeEnvironmentsClient", "Update", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubeenvironments.KubeEnvironmentsClient", "Update", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kubeenvironments.KubeEnvironmentsClient", "Update", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdate prepares the Update request.
func (c KubeEnvironmentsClient) preparerForUpdate(ctx context.Context, id KubeEnvironmentId, input KubeEnvironmentPatchResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdate handles the response to the Update request. The method always
// closes the http.Response Body.
func (c KubeEnvironmentsClient) responderForUpdate(resp *http.Response) (result UpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

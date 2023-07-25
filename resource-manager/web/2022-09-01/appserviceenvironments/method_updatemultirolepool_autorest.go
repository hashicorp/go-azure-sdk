package appserviceenvironments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateMultiRolePoolOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkerPoolResource
}

// UpdateMultiRolePool ...
func (c AppServiceEnvironmentsClient) UpdateMultiRolePool(ctx context.Context, id HostingEnvironmentId, input WorkerPoolResource) (result UpdateMultiRolePoolOperationResponse, err error) {
	req, err := c.preparerForUpdateMultiRolePool(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "UpdateMultiRolePool", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "UpdateMultiRolePool", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateMultiRolePool(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "UpdateMultiRolePool", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateMultiRolePool prepares the UpdateMultiRolePool request.
func (c AppServiceEnvironmentsClient) preparerForUpdateMultiRolePool(ctx context.Context, id HostingEnvironmentId, input WorkerPoolResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/multiRolePools/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateMultiRolePool handles the response to the UpdateMultiRolePool request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForUpdateMultiRolePool(resp *http.Response) (result UpdateMultiRolePoolOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

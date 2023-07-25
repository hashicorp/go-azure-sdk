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

type GetMultiRolePoolOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkerPoolResource
}

// GetMultiRolePool ...
func (c AppServiceEnvironmentsClient) GetMultiRolePool(ctx context.Context, id HostingEnvironmentId) (result GetMultiRolePoolOperationResponse, err error) {
	req, err := c.preparerForGetMultiRolePool(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetMultiRolePool", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetMultiRolePool", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetMultiRolePool(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetMultiRolePool", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetMultiRolePool prepares the GetMultiRolePool request.
func (c AppServiceEnvironmentsClient) preparerForGetMultiRolePool(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/multiRolePools/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetMultiRolePool handles the response to the GetMultiRolePool request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForGetMultiRolePool(resp *http.Response) (result GetMultiRolePoolOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

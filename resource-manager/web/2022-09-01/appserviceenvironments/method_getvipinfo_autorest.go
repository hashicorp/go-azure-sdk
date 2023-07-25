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

type GetVipInfoOperationResponse struct {
	HttpResponse *http.Response
	Model        *AddressResponse
}

// GetVipInfo ...
func (c AppServiceEnvironmentsClient) GetVipInfo(ctx context.Context, id HostingEnvironmentId) (result GetVipInfoOperationResponse, err error) {
	req, err := c.preparerForGetVipInfo(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetVipInfo", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetVipInfo", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetVipInfo(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetVipInfo", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetVipInfo prepares the GetVipInfo request.
func (c AppServiceEnvironmentsClient) preparerForGetVipInfo(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/capacities/virtualip", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetVipInfo handles the response to the GetVipInfo request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForGetVipInfo(resp *http.Response) (result GetVipInfoOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

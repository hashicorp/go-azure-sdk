package vcenter

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVCenterOperationResponse struct {
	HttpResponse *http.Response
	Model        *VCenter
}

// GetVCenter ...
func (c VCenterClient) GetVCenter(ctx context.Context, id VCenterId) (result GetVCenterOperationResponse, err error) {
	req, err := c.preparerForGetVCenter(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "GetVCenter", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "GetVCenter", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetVCenter(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "GetVCenter", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetVCenter prepares the GetVCenter request.
func (c VCenterClient) preparerForGetVCenter(ctx context.Context, id VCenterId) (*http.Request, error) {
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

// responderForGetVCenter handles the response to the GetVCenter request. The method always
// closes the http.Response Body.
func (c VCenterClient) responderForGetVCenter(resp *http.Response) (result GetVCenterOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

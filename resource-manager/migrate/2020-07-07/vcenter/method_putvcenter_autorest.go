package vcenter

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PutVCenterOperationResponse struct {
	HttpResponse *http.Response
}

// PutVCenter ...
func (c VCenterClient) PutVCenter(ctx context.Context, id VCenterId, input VCenter) (result PutVCenterOperationResponse, err error) {
	req, err := c.preparerForPutVCenter(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "PutVCenter", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "PutVCenter", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPutVCenter(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "PutVCenter", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPutVCenter prepares the PutVCenter request.
func (c VCenterClient) preparerForPutVCenter(ctx context.Context, id VCenterId, input VCenter) (*http.Request, error) {
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

// responderForPutVCenter handles the response to the PutVCenter request. The method always
// closes the http.Response Body.
func (c VCenterClient) responderForPutVCenter(resp *http.Response) (result PutVCenterOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package vcenter

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteVCenterOperationResponse struct {
	HttpResponse *http.Response
}

// DeleteVCenter ...
func (c VCenterClient) DeleteVCenter(ctx context.Context, id VCenterId) (result DeleteVCenterOperationResponse, err error) {
	req, err := c.preparerForDeleteVCenter(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "DeleteVCenter", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "DeleteVCenter", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDeleteVCenter(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "vcenter.VCenterClient", "DeleteVCenter", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDeleteVCenter prepares the DeleteVCenter request.
func (c VCenterClient) preparerForDeleteVCenter(ctx context.Context, id VCenterId) (*http.Request, error) {
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

// responderForDeleteVCenter handles the response to the DeleteVCenter request. The method always
// closes the http.Response Body.
func (c VCenterClient) responderForDeleteVCenter(resp *http.Response) (result DeleteVCenterOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

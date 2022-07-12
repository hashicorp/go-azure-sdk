package virtualmachineimages

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgeZoneGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *VirtualMachineImage
}

// EdgeZoneGet ...
func (c VirtualMachineImagesClient) EdgeZoneGet(ctx context.Context, id OfferSkuVersionId) (result EdgeZoneGetOperationResponse, err error) {
	req, err := c.preparerForEdgeZoneGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForEdgeZoneGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForEdgeZoneGet prepares the EdgeZoneGet request.
func (c VirtualMachineImagesClient) preparerForEdgeZoneGet(ctx context.Context, id OfferSkuVersionId) (*http.Request, error) {
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

// responderForEdgeZoneGet handles the response to the EdgeZoneGet request. The method always
// closes the http.Response Body.
func (c VirtualMachineImagesClient) responderForEdgeZoneGet(resp *http.Response) (result EdgeZoneGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

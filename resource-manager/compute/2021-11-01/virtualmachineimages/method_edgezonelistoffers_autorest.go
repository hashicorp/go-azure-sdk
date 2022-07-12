package virtualmachineimages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgeZoneListOffersOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VirtualMachineImageResource
}

// EdgeZoneListOffers ...
func (c VirtualMachineImagesClient) EdgeZoneListOffers(ctx context.Context, id EdgeZonePublisherId) (result EdgeZoneListOffersOperationResponse, err error) {
	req, err := c.preparerForEdgeZoneListOffers(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneListOffers", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneListOffers", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForEdgeZoneListOffers(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneListOffers", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForEdgeZoneListOffers prepares the EdgeZoneListOffers request.
func (c VirtualMachineImagesClient) preparerForEdgeZoneListOffers(ctx context.Context, id EdgeZonePublisherId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/artifactTypes/vmImage/offers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForEdgeZoneListOffers handles the response to the EdgeZoneListOffers request. The method always
// closes the http.Response Body.
func (c VirtualMachineImagesClient) responderForEdgeZoneListOffers(resp *http.Response) (result EdgeZoneListOffersOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

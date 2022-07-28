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

type ListOffersOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VirtualMachineImageResource
}

// ListOffers ...
func (c VirtualMachineImagesClient) ListOffers(ctx context.Context, id PublisherId) (result ListOffersOperationResponse, err error) {
	req, err := c.preparerForListOffers(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "ListOffers", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "ListOffers", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListOffers(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "ListOffers", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListOffers prepares the ListOffers request.
func (c VirtualMachineImagesClient) preparerForListOffers(ctx context.Context, id PublisherId) (*http.Request, error) {
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

// responderForListOffers handles the response to the ListOffers request. The method always
// closes the http.Response Body.
func (c VirtualMachineImagesClient) responderForListOffers(resp *http.Response) (result ListOffersOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

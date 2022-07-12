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

type EdgeZoneListPublishersOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VirtualMachineImageResource
}

// EdgeZoneListPublishers ...
func (c VirtualMachineImagesClient) EdgeZoneListPublishers(ctx context.Context, id EdgeZoneId) (result EdgeZoneListPublishersOperationResponse, err error) {
	req, err := c.preparerForEdgeZoneListPublishers(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneListPublishers", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneListPublishers", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForEdgeZoneListPublishers(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneListPublishers", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForEdgeZoneListPublishers prepares the EdgeZoneListPublishers request.
func (c VirtualMachineImagesClient) preparerForEdgeZoneListPublishers(ctx context.Context, id EdgeZoneId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/publishers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForEdgeZoneListPublishers handles the response to the EdgeZoneListPublishers request. The method always
// closes the http.Response Body.
func (c VirtualMachineImagesClient) responderForEdgeZoneListPublishers(resp *http.Response) (result EdgeZoneListPublishersOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

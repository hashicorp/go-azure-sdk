package resourcehealthmetadata

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetBySiteSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *ResourceHealthMetadata
}

// GetBySiteSlot ...
func (c ResourceHealthMetadataClient) GetBySiteSlot(ctx context.Context, id SlotId) (result GetBySiteSlotOperationResponse, err error) {
	req, err := c.preparerForGetBySiteSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "GetBySiteSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "GetBySiteSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetBySiteSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "GetBySiteSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetBySiteSlot prepares the GetBySiteSlot request.
func (c ResourceHealthMetadataClient) preparerForGetBySiteSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/resourceHealthMetadata/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetBySiteSlot handles the response to the GetBySiteSlot request. The method always
// closes the http.Response Body.
func (c ResourceHealthMetadataClient) responderForGetBySiteSlot(resp *http.Response) (result GetBySiteSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

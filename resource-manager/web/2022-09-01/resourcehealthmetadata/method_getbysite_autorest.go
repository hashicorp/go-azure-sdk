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

type GetBySiteOperationResponse struct {
	HttpResponse *http.Response
	Model        *ResourceHealthMetadata
}

// GetBySite ...
func (c ResourceHealthMetadataClient) GetBySite(ctx context.Context, id SiteId) (result GetBySiteOperationResponse, err error) {
	req, err := c.preparerForGetBySite(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "GetBySite", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "GetBySite", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetBySite(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealthmetadata.ResourceHealthMetadataClient", "GetBySite", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetBySite prepares the GetBySite request.
func (c ResourceHealthMetadataClient) preparerForGetBySite(ctx context.Context, id SiteId) (*http.Request, error) {
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

// responderForGetBySite handles the response to the GetBySite request. The method always
// closes the http.Response Body.
func (c ResourceHealthMetadataClient) responderForGetBySite(resp *http.Response) (result GetBySiteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

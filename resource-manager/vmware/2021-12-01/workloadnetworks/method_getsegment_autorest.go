package workloadnetworks

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSegmentOperationResponse struct {
	HttpResponse *http.Response
	Model        *WorkloadNetworkSegment
}

// GetSegment ...
func (c WorkloadNetworksClient) GetSegment(ctx context.Context, id SegmentId) (result GetSegmentOperationResponse, err error) {
	req, err := c.preparerForGetSegment(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetSegment", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetSegment", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetSegment(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "GetSegment", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetSegment prepares the GetSegment request.
func (c WorkloadNetworksClient) preparerForGetSegment(ctx context.Context, id SegmentId) (*http.Request, error) {
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

// responderForGetSegment handles the response to the GetSegment request. The method always
// closes the http.Response Body.
func (c WorkloadNetworksClient) responderForGetSegment(resp *http.Response) (result GetSegmentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package workloadnetworks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSegmentsOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateSegments ...
func (c WorkloadNetworksClient) CreateSegments(ctx context.Context, id SegmentId, input WorkloadNetworkSegment) (result CreateSegmentsOperationResponse, err error) {
	req, err := c.preparerForCreateSegments(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "CreateSegments", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateSegments(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "CreateSegments", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateSegmentsThenPoll performs CreateSegments then polls until it's completed
func (c WorkloadNetworksClient) CreateSegmentsThenPoll(ctx context.Context, id SegmentId, input WorkloadNetworkSegment) error {
	result, err := c.CreateSegments(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateSegments: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateSegments: %+v", err)
	}

	return nil
}

// preparerForCreateSegments prepares the CreateSegments request.
func (c WorkloadNetworksClient) preparerForCreateSegments(ctx context.Context, id SegmentId, input WorkloadNetworkSegment) (*http.Request, error) {
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

// senderForCreateSegments sends the CreateSegments request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForCreateSegments(ctx context.Context, req *http.Request) (future CreateSegmentsOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

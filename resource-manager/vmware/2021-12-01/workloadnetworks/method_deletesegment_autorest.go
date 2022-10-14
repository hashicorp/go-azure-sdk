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

type DeleteSegmentOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeleteSegment ...
func (c WorkloadNetworksClient) DeleteSegment(ctx context.Context, id SegmentId) (result DeleteSegmentOperationResponse, err error) {
	req, err := c.preparerForDeleteSegment(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "DeleteSegment", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteSegment(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workloadnetworks.WorkloadNetworksClient", "DeleteSegment", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteSegmentThenPoll performs DeleteSegment then polls until it's completed
func (c WorkloadNetworksClient) DeleteSegmentThenPoll(ctx context.Context, id SegmentId) error {
	result, err := c.DeleteSegment(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteSegment: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteSegment: %+v", err)
	}

	return nil
}

// preparerForDeleteSegment prepares the DeleteSegment request.
func (c WorkloadNetworksClient) preparerForDeleteSegment(ctx context.Context, id SegmentId) (*http.Request, error) {
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

// senderForDeleteSegment sends the DeleteSegment request. The method will close the
// http.Response Body if it receives an error.
func (c WorkloadNetworksClient) senderForDeleteSegment(ctx context.Context, req *http.Request) (future DeleteSegmentOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

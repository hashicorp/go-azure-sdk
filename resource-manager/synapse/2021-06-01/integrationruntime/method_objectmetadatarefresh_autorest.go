package integrationruntime

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

type ObjectMetadataRefreshOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ObjectMetadataRefresh ...
func (c IntegrationRuntimeClient) ObjectMetadataRefresh(ctx context.Context, id IntegrationRuntimeId) (result ObjectMetadataRefreshOperationResponse, err error) {
	req, err := c.preparerForObjectMetadataRefresh(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "ObjectMetadataRefresh", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForObjectMetadataRefresh(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "ObjectMetadataRefresh", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ObjectMetadataRefreshThenPoll performs ObjectMetadataRefresh then polls until it's completed
func (c IntegrationRuntimeClient) ObjectMetadataRefreshThenPoll(ctx context.Context, id IntegrationRuntimeId) error {
	result, err := c.ObjectMetadataRefresh(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ObjectMetadataRefresh: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ObjectMetadataRefresh: %+v", err)
	}

	return nil
}

// preparerForObjectMetadataRefresh prepares the ObjectMetadataRefresh request.
func (c IntegrationRuntimeClient) preparerForObjectMetadataRefresh(ctx context.Context, id IntegrationRuntimeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/refreshObjectMetadata", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForObjectMetadataRefresh sends the ObjectMetadataRefresh request. The method will close the
// http.Response Body if it receives an error.
func (c IntegrationRuntimeClient) senderForObjectMetadataRefresh(ctx context.Context, req *http.Request) (future ObjectMetadataRefreshOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

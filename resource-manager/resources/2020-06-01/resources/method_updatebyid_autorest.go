package resources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateByIdOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// UpdateById ...
func (c ResourcesClient) UpdateById(ctx context.Context, id commonids.ScopeId, input GenericResource) (result UpdateByIdOperationResponse, err error) {
	req, err := c.preparerForUpdateById(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ResourcesClient", "UpdateById", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForUpdateById(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ResourcesClient", "UpdateById", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// UpdateByIdThenPoll performs UpdateById then polls until it's completed
func (c ResourcesClient) UpdateByIdThenPoll(ctx context.Context, id commonids.ScopeId, input GenericResource) error {
	result, err := c.UpdateById(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing UpdateById: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after UpdateById: %+v", err)
	}

	return nil
}

// preparerForUpdateById prepares the UpdateById request.
func (c ResourcesClient) preparerForUpdateById(ctx context.Context, id commonids.ScopeId, input GenericResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForUpdateById sends the UpdateById request. The method will close the
// http.Response Body if it receives an error.
func (c ResourcesClient) senderForUpdateById(ctx context.Context, req *http.Request) (future UpdateByIdOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}

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

type DeleteByIdOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeleteById ...
func (c ResourcesClient) DeleteById(ctx context.Context, id commonids.ScopeId) (result DeleteByIdOperationResponse, err error) {
	req, err := c.preparerForDeleteById(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ResourcesClient", "DeleteById", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteById(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resources.ResourcesClient", "DeleteById", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteByIdThenPoll performs DeleteById then polls until it's completed
func (c ResourcesClient) DeleteByIdThenPoll(ctx context.Context, id commonids.ScopeId) error {
	result, err := c.DeleteById(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteById: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteById: %+v", err)
	}

	return nil
}

// preparerForDeleteById prepares the DeleteById request.
func (c ResourcesClient) preparerForDeleteById(ctx context.Context, id commonids.ScopeId) (*http.Request, error) {
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

// senderForDeleteById sends the DeleteById request. The method will close the
// http.Response Body if it receives an error.
func (c ResourcesClient) senderForDeleteById(ctx context.Context, req *http.Request) (future DeleteByIdOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

package trustedaccess

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

type RoleBindingsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RoleBindingsDelete ...
func (c TrustedAccessClient) RoleBindingsDelete(ctx context.Context, id TrustedAccessRoleBindingId) (result RoleBindingsDeleteOperationResponse, err error) {
	req, err := c.preparerForRoleBindingsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trustedaccess.TrustedAccessClient", "RoleBindingsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRoleBindingsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trustedaccess.TrustedAccessClient", "RoleBindingsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RoleBindingsDeleteThenPoll performs RoleBindingsDelete then polls until it's completed
func (c TrustedAccessClient) RoleBindingsDeleteThenPoll(ctx context.Context, id TrustedAccessRoleBindingId) error {
	result, err := c.RoleBindingsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RoleBindingsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RoleBindingsDelete: %+v", err)
	}

	return nil
}

// preparerForRoleBindingsDelete prepares the RoleBindingsDelete request.
func (c TrustedAccessClient) preparerForRoleBindingsDelete(ctx context.Context, id TrustedAccessRoleBindingId) (*http.Request, error) {
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

// senderForRoleBindingsDelete sends the RoleBindingsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c TrustedAccessClient) senderForRoleBindingsDelete(ctx context.Context, req *http.Request) (future RoleBindingsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

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

type RoleBindingsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RoleBindingsCreateOrUpdate ...
func (c TrustedAccessClient) RoleBindingsCreateOrUpdate(ctx context.Context, id TrustedAccessRoleBindingId, input TrustedAccessRoleBinding) (result RoleBindingsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForRoleBindingsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trustedaccess.TrustedAccessClient", "RoleBindingsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRoleBindingsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trustedaccess.TrustedAccessClient", "RoleBindingsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RoleBindingsCreateOrUpdateThenPoll performs RoleBindingsCreateOrUpdate then polls until it's completed
func (c TrustedAccessClient) RoleBindingsCreateOrUpdateThenPoll(ctx context.Context, id TrustedAccessRoleBindingId, input TrustedAccessRoleBinding) error {
	result, err := c.RoleBindingsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RoleBindingsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RoleBindingsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForRoleBindingsCreateOrUpdate prepares the RoleBindingsCreateOrUpdate request.
func (c TrustedAccessClient) preparerForRoleBindingsCreateOrUpdate(ctx context.Context, id TrustedAccessRoleBindingId, input TrustedAccessRoleBinding) (*http.Request, error) {
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

// senderForRoleBindingsCreateOrUpdate sends the RoleBindingsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c TrustedAccessClient) senderForRoleBindingsCreateOrUpdate(ctx context.Context, req *http.Request) (future RoleBindingsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

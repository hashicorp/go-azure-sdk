package deployments

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

type ValidateAtManagementGroupScopeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ValidateAtManagementGroupScope ...
func (c DeploymentsClient) ValidateAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId, input ScopedDeployment) (result ValidateAtManagementGroupScopeOperationResponse, err error) {
	req, err := c.preparerForValidateAtManagementGroupScope(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ValidateAtManagementGroupScope", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForValidateAtManagementGroupScope(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ValidateAtManagementGroupScope", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ValidateAtManagementGroupScopeThenPoll performs ValidateAtManagementGroupScope then polls until it's completed
func (c DeploymentsClient) ValidateAtManagementGroupScopeThenPoll(ctx context.Context, id Providers2DeploymentId, input ScopedDeployment) error {
	result, err := c.ValidateAtManagementGroupScope(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ValidateAtManagementGroupScope: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ValidateAtManagementGroupScope: %+v", err)
	}

	return nil
}

// preparerForValidateAtManagementGroupScope prepares the ValidateAtManagementGroupScope request.
func (c DeploymentsClient) preparerForValidateAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId, input ScopedDeployment) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/validate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForValidateAtManagementGroupScope sends the ValidateAtManagementGroupScope request. The method will close the
// http.Response Body if it receives an error.
func (c DeploymentsClient) senderForValidateAtManagementGroupScope(ctx context.Context, req *http.Request) (future ValidateAtManagementGroupScopeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

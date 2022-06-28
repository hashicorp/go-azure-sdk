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

type CreateOrUpdateAtManagementGroupScopeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdateAtManagementGroupScope ...
func (c DeploymentsClient) CreateOrUpdateAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId, input ScopedDeployment) (result CreateOrUpdateAtManagementGroupScopeOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateAtManagementGroupScope(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CreateOrUpdateAtManagementGroupScope", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdateAtManagementGroupScope(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CreateOrUpdateAtManagementGroupScope", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateAtManagementGroupScopeThenPoll performs CreateOrUpdateAtManagementGroupScope then polls until it's completed
func (c DeploymentsClient) CreateOrUpdateAtManagementGroupScopeThenPoll(ctx context.Context, id Providers2DeploymentId, input ScopedDeployment) error {
	result, err := c.CreateOrUpdateAtManagementGroupScope(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateAtManagementGroupScope: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateAtManagementGroupScope: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdateAtManagementGroupScope prepares the CreateOrUpdateAtManagementGroupScope request.
func (c DeploymentsClient) preparerForCreateOrUpdateAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId, input ScopedDeployment) (*http.Request, error) {
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

// senderForCreateOrUpdateAtManagementGroupScope sends the CreateOrUpdateAtManagementGroupScope request. The method will close the
// http.Response Body if it receives an error.
func (c DeploymentsClient) senderForCreateOrUpdateAtManagementGroupScope(ctx context.Context, req *http.Request) (future CreateOrUpdateAtManagementGroupScopeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}

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

type CreateOrUpdateAtTenantScopeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdateAtTenantScope ...
func (c DeploymentsClient) CreateOrUpdateAtTenantScope(ctx context.Context, id DeploymentId, input ScopedDeployment) (result CreateOrUpdateAtTenantScopeOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateAtTenantScope(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CreateOrUpdateAtTenantScope", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdateAtTenantScope(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CreateOrUpdateAtTenantScope", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateAtTenantScopeThenPoll performs CreateOrUpdateAtTenantScope then polls until it's completed
func (c DeploymentsClient) CreateOrUpdateAtTenantScopeThenPoll(ctx context.Context, id DeploymentId, input ScopedDeployment) error {
	result, err := c.CreateOrUpdateAtTenantScope(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateAtTenantScope: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateAtTenantScope: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdateAtTenantScope prepares the CreateOrUpdateAtTenantScope request.
func (c DeploymentsClient) preparerForCreateOrUpdateAtTenantScope(ctx context.Context, id DeploymentId, input ScopedDeployment) (*http.Request, error) {
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

// senderForCreateOrUpdateAtTenantScope sends the CreateOrUpdateAtTenantScope request. The method will close the
// http.Response Body if it receives an error.
func (c DeploymentsClient) senderForCreateOrUpdateAtTenantScope(ctx context.Context, req *http.Request) (future CreateOrUpdateAtTenantScopeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

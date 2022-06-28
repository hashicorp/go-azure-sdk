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

type ValidateAtTenantScopeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ValidateAtTenantScope ...
func (c DeploymentsClient) ValidateAtTenantScope(ctx context.Context, id DeploymentId, input ScopedDeployment) (result ValidateAtTenantScopeOperationResponse, err error) {
	req, err := c.preparerForValidateAtTenantScope(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ValidateAtTenantScope", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForValidateAtTenantScope(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ValidateAtTenantScope", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ValidateAtTenantScopeThenPoll performs ValidateAtTenantScope then polls until it's completed
func (c DeploymentsClient) ValidateAtTenantScopeThenPoll(ctx context.Context, id DeploymentId, input ScopedDeployment) error {
	result, err := c.ValidateAtTenantScope(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ValidateAtTenantScope: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ValidateAtTenantScope: %+v", err)
	}

	return nil
}

// preparerForValidateAtTenantScope prepares the ValidateAtTenantScope request.
func (c DeploymentsClient) preparerForValidateAtTenantScope(ctx context.Context, id DeploymentId, input ScopedDeployment) (*http.Request, error) {
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

// senderForValidateAtTenantScope sends the ValidateAtTenantScope request. The method will close the
// http.Response Body if it receives an error.
func (c DeploymentsClient) senderForValidateAtTenantScope(ctx context.Context, req *http.Request) (future ValidateAtTenantScopeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}

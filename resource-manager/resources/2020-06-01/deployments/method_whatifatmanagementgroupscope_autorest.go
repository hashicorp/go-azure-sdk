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

type WhatIfAtManagementGroupScopeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WhatIfAtManagementGroupScope ...
func (c DeploymentsClient) WhatIfAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId, input ScopedDeploymentWhatIf) (result WhatIfAtManagementGroupScopeOperationResponse, err error) {
	req, err := c.preparerForWhatIfAtManagementGroupScope(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "WhatIfAtManagementGroupScope", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWhatIfAtManagementGroupScope(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "WhatIfAtManagementGroupScope", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WhatIfAtManagementGroupScopeThenPoll performs WhatIfAtManagementGroupScope then polls until it's completed
func (c DeploymentsClient) WhatIfAtManagementGroupScopeThenPoll(ctx context.Context, id Providers2DeploymentId, input ScopedDeploymentWhatIf) error {
	result, err := c.WhatIfAtManagementGroupScope(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WhatIfAtManagementGroupScope: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WhatIfAtManagementGroupScope: %+v", err)
	}

	return nil
}

// preparerForWhatIfAtManagementGroupScope prepares the WhatIfAtManagementGroupScope request.
func (c DeploymentsClient) preparerForWhatIfAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId, input ScopedDeploymentWhatIf) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/whatIf", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForWhatIfAtManagementGroupScope sends the WhatIfAtManagementGroupScope request. The method will close the
// http.Response Body if it receives an error.
func (c DeploymentsClient) senderForWhatIfAtManagementGroupScope(ctx context.Context, req *http.Request) (future WhatIfAtManagementGroupScopeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

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

type WhatIfAtTenantScopeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WhatIfAtTenantScope ...
func (c DeploymentsClient) WhatIfAtTenantScope(ctx context.Context, id DeploymentId, input ScopedDeploymentWhatIf) (result WhatIfAtTenantScopeOperationResponse, err error) {
	req, err := c.preparerForWhatIfAtTenantScope(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "WhatIfAtTenantScope", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWhatIfAtTenantScope(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "WhatIfAtTenantScope", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WhatIfAtTenantScopeThenPoll performs WhatIfAtTenantScope then polls until it's completed
func (c DeploymentsClient) WhatIfAtTenantScopeThenPoll(ctx context.Context, id DeploymentId, input ScopedDeploymentWhatIf) error {
	result, err := c.WhatIfAtTenantScope(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WhatIfAtTenantScope: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WhatIfAtTenantScope: %+v", err)
	}

	return nil
}

// preparerForWhatIfAtTenantScope prepares the WhatIfAtTenantScope request.
func (c DeploymentsClient) preparerForWhatIfAtTenantScope(ctx context.Context, id DeploymentId, input ScopedDeploymentWhatIf) (*http.Request, error) {
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

// senderForWhatIfAtTenantScope sends the WhatIfAtTenantScope request. The method will close the
// http.Response Body if it receives an error.
func (c DeploymentsClient) senderForWhatIfAtTenantScope(ctx context.Context, req *http.Request) (future WhatIfAtTenantScopeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

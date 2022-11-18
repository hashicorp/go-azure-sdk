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

type WhatIfOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// WhatIf ...
func (c DeploymentsClient) WhatIf(ctx context.Context, id ResourceGroupProviderDeploymentId, input DeploymentWhatIf) (result WhatIfOperationResponse, err error) {
	req, err := c.preparerForWhatIf(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "WhatIf", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForWhatIf(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "WhatIf", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// WhatIfThenPoll performs WhatIf then polls until it's completed
func (c DeploymentsClient) WhatIfThenPoll(ctx context.Context, id ResourceGroupProviderDeploymentId, input DeploymentWhatIf) error {
	result, err := c.WhatIf(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WhatIf: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after WhatIf: %+v", err)
	}

	return nil
}

// preparerForWhatIf prepares the WhatIf request.
func (c DeploymentsClient) preparerForWhatIf(ctx context.Context, id ResourceGroupProviderDeploymentId, input DeploymentWhatIf) (*http.Request, error) {
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

// senderForWhatIf sends the WhatIf request. The method will close the
// http.Response Body if it receives an error.
func (c DeploymentsClient) senderForWhatIf(ctx context.Context, req *http.Request) (future WhatIfOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

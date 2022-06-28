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

type DeleteAtSubscriptionScopeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeleteAtSubscriptionScope ...
func (c DeploymentsClient) DeleteAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId) (result DeleteAtSubscriptionScopeOperationResponse, err error) {
	req, err := c.preparerForDeleteAtSubscriptionScope(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "DeleteAtSubscriptionScope", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteAtSubscriptionScope(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "DeleteAtSubscriptionScope", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteAtSubscriptionScopeThenPoll performs DeleteAtSubscriptionScope then polls until it's completed
func (c DeploymentsClient) DeleteAtSubscriptionScopeThenPoll(ctx context.Context, id ProviderDeploymentId) error {
	result, err := c.DeleteAtSubscriptionScope(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteAtSubscriptionScope: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteAtSubscriptionScope: %+v", err)
	}

	return nil
}

// preparerForDeleteAtSubscriptionScope prepares the DeleteAtSubscriptionScope request.
func (c DeploymentsClient) preparerForDeleteAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId) (*http.Request, error) {
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

// senderForDeleteAtSubscriptionScope sends the DeleteAtSubscriptionScope request. The method will close the
// http.Response Body if it receives an error.
func (c DeploymentsClient) senderForDeleteAtSubscriptionScope(ctx context.Context, req *http.Request) (future DeleteAtSubscriptionScopeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}

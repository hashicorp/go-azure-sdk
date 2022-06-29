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

type CreateOrUpdateAtSubscriptionScopeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdateAtSubscriptionScope ...
func (c DeploymentsClient) CreateOrUpdateAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId, input Deployment) (result CreateOrUpdateAtSubscriptionScopeOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateAtSubscriptionScope(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CreateOrUpdateAtSubscriptionScope", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdateAtSubscriptionScope(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CreateOrUpdateAtSubscriptionScope", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateAtSubscriptionScopeThenPoll performs CreateOrUpdateAtSubscriptionScope then polls until it's completed
func (c DeploymentsClient) CreateOrUpdateAtSubscriptionScopeThenPoll(ctx context.Context, id ProviderDeploymentId, input Deployment) error {
	result, err := c.CreateOrUpdateAtSubscriptionScope(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateAtSubscriptionScope: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateAtSubscriptionScope: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdateAtSubscriptionScope prepares the CreateOrUpdateAtSubscriptionScope request.
func (c DeploymentsClient) preparerForCreateOrUpdateAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId, input Deployment) (*http.Request, error) {
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

// senderForCreateOrUpdateAtSubscriptionScope sends the CreateOrUpdateAtSubscriptionScope request. The method will close the
// http.Response Body if it receives an error.
func (c DeploymentsClient) senderForCreateOrUpdateAtSubscriptionScope(ctx context.Context, req *http.Request) (future CreateOrUpdateAtSubscriptionScopeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

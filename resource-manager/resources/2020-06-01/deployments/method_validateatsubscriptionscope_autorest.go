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

type ValidateAtSubscriptionScopeOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ValidateAtSubscriptionScope ...
func (c DeploymentsClient) ValidateAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId, input Deployment) (result ValidateAtSubscriptionScopeOperationResponse, err error) {
	req, err := c.preparerForValidateAtSubscriptionScope(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ValidateAtSubscriptionScope", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForValidateAtSubscriptionScope(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "ValidateAtSubscriptionScope", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ValidateAtSubscriptionScopeThenPoll performs ValidateAtSubscriptionScope then polls until it's completed
func (c DeploymentsClient) ValidateAtSubscriptionScopeThenPoll(ctx context.Context, id ProviderDeploymentId, input Deployment) error {
	result, err := c.ValidateAtSubscriptionScope(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ValidateAtSubscriptionScope: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ValidateAtSubscriptionScope: %+v", err)
	}

	return nil
}

// preparerForValidateAtSubscriptionScope prepares the ValidateAtSubscriptionScope request.
func (c DeploymentsClient) preparerForValidateAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId, input Deployment) (*http.Request, error) {
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

// senderForValidateAtSubscriptionScope sends the ValidateAtSubscriptionScope request. The method will close the
// http.Response Body if it receives an error.
func (c DeploymentsClient) senderForValidateAtSubscriptionScope(ctx context.Context, req *http.Request) (future ValidateAtSubscriptionScopeOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewLongRunningPollerFromResponse(ctx, resp, c.Client)
	return
}

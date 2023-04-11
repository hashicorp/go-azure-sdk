package adaptivenetworkhardenings

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

type EnforceOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// Enforce ...
func (c AdaptiveNetworkHardeningsClient) Enforce(ctx context.Context, id AdaptiveNetworkHardeningId, input AdaptiveNetworkHardeningEnforceRequest) (result EnforceOperationResponse, err error) {
	req, err := c.preparerForEnforce(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adaptivenetworkhardenings.AdaptiveNetworkHardeningsClient", "Enforce", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForEnforce(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "adaptivenetworkhardenings.AdaptiveNetworkHardeningsClient", "Enforce", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// EnforceThenPoll performs Enforce then polls until it's completed
func (c AdaptiveNetworkHardeningsClient) EnforceThenPoll(ctx context.Context, id AdaptiveNetworkHardeningId, input AdaptiveNetworkHardeningEnforceRequest) error {
	result, err := c.Enforce(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing Enforce: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Enforce: %+v", err)
	}

	return nil
}

// preparerForEnforce prepares the Enforce request.
func (c AdaptiveNetworkHardeningsClient) preparerForEnforce(ctx context.Context, id AdaptiveNetworkHardeningId, input AdaptiveNetworkHardeningEnforceRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/enforce", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForEnforce sends the Enforce request. The method will close the
// http.Response Body if it receives an error.
func (c AdaptiveNetworkHardeningsClient) senderForEnforce(ctx context.Context, req *http.Request) (future EnforceOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

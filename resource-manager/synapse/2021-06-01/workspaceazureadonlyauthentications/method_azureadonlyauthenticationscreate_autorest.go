package workspaceazureadonlyauthentications

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

type AzureADOnlyAuthenticationsCreateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// AzureADOnlyAuthenticationsCreate ...
func (c WorkspaceAzureADOnlyAuthenticationsClient) AzureADOnlyAuthenticationsCreate(ctx context.Context, id WorkspaceId, input AzureADOnlyAuthentication) (result AzureADOnlyAuthenticationsCreateOperationResponse, err error) {
	req, err := c.preparerForAzureADOnlyAuthenticationsCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaceazureadonlyauthentications.WorkspaceAzureADOnlyAuthenticationsClient", "AzureADOnlyAuthenticationsCreate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForAzureADOnlyAuthenticationsCreate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaceazureadonlyauthentications.WorkspaceAzureADOnlyAuthenticationsClient", "AzureADOnlyAuthenticationsCreate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// AzureADOnlyAuthenticationsCreateThenPoll performs AzureADOnlyAuthenticationsCreate then polls until it's completed
func (c WorkspaceAzureADOnlyAuthenticationsClient) AzureADOnlyAuthenticationsCreateThenPoll(ctx context.Context, id WorkspaceId, input AzureADOnlyAuthentication) error {
	result, err := c.AzureADOnlyAuthenticationsCreate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing AzureADOnlyAuthenticationsCreate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after AzureADOnlyAuthenticationsCreate: %+v", err)
	}

	return nil
}

// preparerForAzureADOnlyAuthenticationsCreate prepares the AzureADOnlyAuthenticationsCreate request.
func (c WorkspaceAzureADOnlyAuthenticationsClient) preparerForAzureADOnlyAuthenticationsCreate(ctx context.Context, id WorkspaceId, input AzureADOnlyAuthentication) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/azureADOnlyAuthentications/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForAzureADOnlyAuthenticationsCreate sends the AzureADOnlyAuthenticationsCreate request. The method will close the
// http.Response Body if it receives an error.
func (c WorkspaceAzureADOnlyAuthenticationsClient) senderForAzureADOnlyAuthenticationsCreate(ctx context.Context, req *http.Request) (future AzureADOnlyAuthenticationsCreateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

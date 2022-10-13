package appplatform

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

type GatewayCustomDomainsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// GatewayCustomDomainsDelete ...
func (c AppPlatformClient) GatewayCustomDomainsDelete(ctx context.Context, id GatewayDomainId) (result GatewayCustomDomainsDeleteOperationResponse, err error) {
	req, err := c.preparerForGatewayCustomDomainsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayCustomDomainsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForGatewayCustomDomainsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewayCustomDomainsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// GatewayCustomDomainsDeleteThenPoll performs GatewayCustomDomainsDelete then polls until it's completed
func (c AppPlatformClient) GatewayCustomDomainsDeleteThenPoll(ctx context.Context, id GatewayDomainId) error {
	result, err := c.GatewayCustomDomainsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing GatewayCustomDomainsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after GatewayCustomDomainsDelete: %+v", err)
	}

	return nil
}

// preparerForGatewayCustomDomainsDelete prepares the GatewayCustomDomainsDelete request.
func (c AppPlatformClient) preparerForGatewayCustomDomainsDelete(ctx context.Context, id GatewayDomainId) (*http.Request, error) {
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

// senderForGatewayCustomDomainsDelete sends the GatewayCustomDomainsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForGatewayCustomDomainsDelete(ctx context.Context, req *http.Request) (future GatewayCustomDomainsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

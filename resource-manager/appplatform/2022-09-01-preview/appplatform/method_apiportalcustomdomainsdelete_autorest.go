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

type ApiPortalCustomDomainsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ApiPortalCustomDomainsDelete ...
func (c AppPlatformClient) ApiPortalCustomDomainsDelete(ctx context.Context, id ApiPortalDomainId) (result ApiPortalCustomDomainsDeleteOperationResponse, err error) {
	req, err := c.preparerForApiPortalCustomDomainsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalCustomDomainsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForApiPortalCustomDomainsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalCustomDomainsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ApiPortalCustomDomainsDeleteThenPoll performs ApiPortalCustomDomainsDelete then polls until it's completed
func (c AppPlatformClient) ApiPortalCustomDomainsDeleteThenPoll(ctx context.Context, id ApiPortalDomainId) error {
	result, err := c.ApiPortalCustomDomainsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ApiPortalCustomDomainsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ApiPortalCustomDomainsDelete: %+v", err)
	}

	return nil
}

// preparerForApiPortalCustomDomainsDelete prepares the ApiPortalCustomDomainsDelete request.
func (c AppPlatformClient) preparerForApiPortalCustomDomainsDelete(ctx context.Context, id ApiPortalDomainId) (*http.Request, error) {
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

// senderForApiPortalCustomDomainsDelete sends the ApiPortalCustomDomainsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForApiPortalCustomDomainsDelete(ctx context.Context, req *http.Request) (future ApiPortalCustomDomainsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

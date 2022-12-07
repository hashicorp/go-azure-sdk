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

type ApiPortalCustomDomainsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ApiPortalCustomDomainsCreateOrUpdate ...
func (c AppPlatformClient) ApiPortalCustomDomainsCreateOrUpdate(ctx context.Context, id ApiPortalDomainId, input ApiPortalCustomDomainResource) (result ApiPortalCustomDomainsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForApiPortalCustomDomainsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalCustomDomainsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForApiPortalCustomDomainsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalCustomDomainsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ApiPortalCustomDomainsCreateOrUpdateThenPoll performs ApiPortalCustomDomainsCreateOrUpdate then polls until it's completed
func (c AppPlatformClient) ApiPortalCustomDomainsCreateOrUpdateThenPoll(ctx context.Context, id ApiPortalDomainId, input ApiPortalCustomDomainResource) error {
	result, err := c.ApiPortalCustomDomainsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ApiPortalCustomDomainsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ApiPortalCustomDomainsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForApiPortalCustomDomainsCreateOrUpdate prepares the ApiPortalCustomDomainsCreateOrUpdate request.
func (c AppPlatformClient) preparerForApiPortalCustomDomainsCreateOrUpdate(ctx context.Context, id ApiPortalDomainId, input ApiPortalCustomDomainResource) (*http.Request, error) {
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

// senderForApiPortalCustomDomainsCreateOrUpdate sends the ApiPortalCustomDomainsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForApiPortalCustomDomainsCreateOrUpdate(ctx context.Context, req *http.Request) (future ApiPortalCustomDomainsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

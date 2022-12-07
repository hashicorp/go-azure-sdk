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

type CustomDomainsUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CustomDomainsUpdate ...
func (c AppPlatformClient) CustomDomainsUpdate(ctx context.Context, id DomainId, input CustomDomainResource) (result CustomDomainsUpdateOperationResponse, err error) {
	req, err := c.preparerForCustomDomainsUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "CustomDomainsUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCustomDomainsUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "CustomDomainsUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CustomDomainsUpdateThenPoll performs CustomDomainsUpdate then polls until it's completed
func (c AppPlatformClient) CustomDomainsUpdateThenPoll(ctx context.Context, id DomainId, input CustomDomainResource) error {
	result, err := c.CustomDomainsUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CustomDomainsUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CustomDomainsUpdate: %+v", err)
	}

	return nil
}

// preparerForCustomDomainsUpdate prepares the CustomDomainsUpdate request.
func (c AppPlatformClient) preparerForCustomDomainsUpdate(ctx context.Context, id DomainId, input CustomDomainResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCustomDomainsUpdate sends the CustomDomainsUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForCustomDomainsUpdate(ctx context.Context, req *http.Request) (future CustomDomainsUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

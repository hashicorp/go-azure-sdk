package staticsites

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

type DeleteStaticSiteCustomDomainOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeleteStaticSiteCustomDomain ...
func (c StaticSitesClient) DeleteStaticSiteCustomDomain(ctx context.Context, id CustomDomainId) (result DeleteStaticSiteCustomDomainOperationResponse, err error) {
	req, err := c.preparerForDeleteStaticSiteCustomDomain(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteStaticSiteCustomDomain", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteStaticSiteCustomDomain(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteStaticSiteCustomDomain", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteStaticSiteCustomDomainThenPoll performs DeleteStaticSiteCustomDomain then polls until it's completed
func (c StaticSitesClient) DeleteStaticSiteCustomDomainThenPoll(ctx context.Context, id CustomDomainId) error {
	result, err := c.DeleteStaticSiteCustomDomain(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteStaticSiteCustomDomain: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteStaticSiteCustomDomain: %+v", err)
	}

	return nil
}

// preparerForDeleteStaticSiteCustomDomain prepares the DeleteStaticSiteCustomDomain request.
func (c StaticSitesClient) preparerForDeleteStaticSiteCustomDomain(ctx context.Context, id CustomDomainId) (*http.Request, error) {
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

// senderForDeleteStaticSiteCustomDomain sends the DeleteStaticSiteCustomDomain request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForDeleteStaticSiteCustomDomain(ctx context.Context, req *http.Request) (future DeleteStaticSiteCustomDomainOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

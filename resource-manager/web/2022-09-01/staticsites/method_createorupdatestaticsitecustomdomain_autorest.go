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

type CreateOrUpdateStaticSiteCustomDomainOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdateStaticSiteCustomDomain ...
func (c StaticSitesClient) CreateOrUpdateStaticSiteCustomDomain(ctx context.Context, id CustomDomainId, input StaticSiteCustomDomainRequestPropertiesARMResource) (result CreateOrUpdateStaticSiteCustomDomainOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateStaticSiteCustomDomain(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteCustomDomain", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdateStaticSiteCustomDomain(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "CreateOrUpdateStaticSiteCustomDomain", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateStaticSiteCustomDomainThenPoll performs CreateOrUpdateStaticSiteCustomDomain then polls until it's completed
func (c StaticSitesClient) CreateOrUpdateStaticSiteCustomDomainThenPoll(ctx context.Context, id CustomDomainId, input StaticSiteCustomDomainRequestPropertiesARMResource) error {
	result, err := c.CreateOrUpdateStaticSiteCustomDomain(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateStaticSiteCustomDomain: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateStaticSiteCustomDomain: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdateStaticSiteCustomDomain prepares the CreateOrUpdateStaticSiteCustomDomain request.
func (c StaticSitesClient) preparerForCreateOrUpdateStaticSiteCustomDomain(ctx context.Context, id CustomDomainId, input StaticSiteCustomDomainRequestPropertiesARMResource) (*http.Request, error) {
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

// senderForCreateOrUpdateStaticSiteCustomDomain sends the CreateOrUpdateStaticSiteCustomDomain request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForCreateOrUpdateStaticSiteCustomDomain(ctx context.Context, req *http.Request) (future CreateOrUpdateStaticSiteCustomDomainOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

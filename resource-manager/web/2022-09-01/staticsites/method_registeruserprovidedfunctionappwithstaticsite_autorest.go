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

type RegisterUserProvidedFunctionAppWithStaticSiteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type RegisterUserProvidedFunctionAppWithStaticSiteOperationOptions struct {
	IsForced *bool
}

func DefaultRegisterUserProvidedFunctionAppWithStaticSiteOperationOptions() RegisterUserProvidedFunctionAppWithStaticSiteOperationOptions {
	return RegisterUserProvidedFunctionAppWithStaticSiteOperationOptions{}
}

func (o RegisterUserProvidedFunctionAppWithStaticSiteOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RegisterUserProvidedFunctionAppWithStaticSiteOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IsForced != nil {
		out["isForced"] = *o.IsForced
	}

	return out
}

// RegisterUserProvidedFunctionAppWithStaticSite ...
func (c StaticSitesClient) RegisterUserProvidedFunctionAppWithStaticSite(ctx context.Context, id UserProvidedFunctionAppId, input StaticSiteUserProvidedFunctionAppARMResource, options RegisterUserProvidedFunctionAppWithStaticSiteOperationOptions) (result RegisterUserProvidedFunctionAppWithStaticSiteOperationResponse, err error) {
	req, err := c.preparerForRegisterUserProvidedFunctionAppWithStaticSite(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "RegisterUserProvidedFunctionAppWithStaticSite", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegisterUserProvidedFunctionAppWithStaticSite(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "RegisterUserProvidedFunctionAppWithStaticSite", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegisterUserProvidedFunctionAppWithStaticSiteThenPoll performs RegisterUserProvidedFunctionAppWithStaticSite then polls until it's completed
func (c StaticSitesClient) RegisterUserProvidedFunctionAppWithStaticSiteThenPoll(ctx context.Context, id UserProvidedFunctionAppId, input StaticSiteUserProvidedFunctionAppARMResource, options RegisterUserProvidedFunctionAppWithStaticSiteOperationOptions) error {
	result, err := c.RegisterUserProvidedFunctionAppWithStaticSite(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing RegisterUserProvidedFunctionAppWithStaticSite: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegisterUserProvidedFunctionAppWithStaticSite: %+v", err)
	}

	return nil
}

// preparerForRegisterUserProvidedFunctionAppWithStaticSite prepares the RegisterUserProvidedFunctionAppWithStaticSite request.
func (c StaticSitesClient) preparerForRegisterUserProvidedFunctionAppWithStaticSite(ctx context.Context, id UserProvidedFunctionAppId, input StaticSiteUserProvidedFunctionAppARMResource, options RegisterUserProvidedFunctionAppWithStaticSiteOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRegisterUserProvidedFunctionAppWithStaticSite sends the RegisterUserProvidedFunctionAppWithStaticSite request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForRegisterUserProvidedFunctionAppWithStaticSite(ctx context.Context, req *http.Request) (future RegisterUserProvidedFunctionAppWithStaticSiteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

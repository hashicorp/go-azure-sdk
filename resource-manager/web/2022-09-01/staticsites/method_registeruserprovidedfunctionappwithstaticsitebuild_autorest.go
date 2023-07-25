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

type RegisterUserProvidedFunctionAppWithStaticSiteBuildOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type RegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions struct {
	IsForced *bool
}

func DefaultRegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions() RegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions {
	return RegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions{}
}

func (o RegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IsForced != nil {
		out["isForced"] = *o.IsForced
	}

	return out
}

// RegisterUserProvidedFunctionAppWithStaticSiteBuild ...
func (c StaticSitesClient) RegisterUserProvidedFunctionAppWithStaticSiteBuild(ctx context.Context, id BuildUserProvidedFunctionAppId, input StaticSiteUserProvidedFunctionAppARMResource, options RegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions) (result RegisterUserProvidedFunctionAppWithStaticSiteBuildOperationResponse, err error) {
	req, err := c.preparerForRegisterUserProvidedFunctionAppWithStaticSiteBuild(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "RegisterUserProvidedFunctionAppWithStaticSiteBuild", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegisterUserProvidedFunctionAppWithStaticSiteBuild(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "RegisterUserProvidedFunctionAppWithStaticSiteBuild", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegisterUserProvidedFunctionAppWithStaticSiteBuildThenPoll performs RegisterUserProvidedFunctionAppWithStaticSiteBuild then polls until it's completed
func (c StaticSitesClient) RegisterUserProvidedFunctionAppWithStaticSiteBuildThenPoll(ctx context.Context, id BuildUserProvidedFunctionAppId, input StaticSiteUserProvidedFunctionAppARMResource, options RegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions) error {
	result, err := c.RegisterUserProvidedFunctionAppWithStaticSiteBuild(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing RegisterUserProvidedFunctionAppWithStaticSiteBuild: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegisterUserProvidedFunctionAppWithStaticSiteBuild: %+v", err)
	}

	return nil
}

// preparerForRegisterUserProvidedFunctionAppWithStaticSiteBuild prepares the RegisterUserProvidedFunctionAppWithStaticSiteBuild request.
func (c StaticSitesClient) preparerForRegisterUserProvidedFunctionAppWithStaticSiteBuild(ctx context.Context, id BuildUserProvidedFunctionAppId, input StaticSiteUserProvidedFunctionAppARMResource, options RegisterUserProvidedFunctionAppWithStaticSiteBuildOperationOptions) (*http.Request, error) {
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

// senderForRegisterUserProvidedFunctionAppWithStaticSiteBuild sends the RegisterUserProvidedFunctionAppWithStaticSiteBuild request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForRegisterUserProvidedFunctionAppWithStaticSiteBuild(ctx context.Context, req *http.Request) (future RegisterUserProvidedFunctionAppWithStaticSiteBuildOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

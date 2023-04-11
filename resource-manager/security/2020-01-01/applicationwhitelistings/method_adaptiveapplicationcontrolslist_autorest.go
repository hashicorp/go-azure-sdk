package applicationwhitelistings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdaptiveApplicationControlsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *AdaptiveApplicationControlGroups
}

type AdaptiveApplicationControlsListOperationOptions struct {
	IncludePathRecommendations *bool
	Summary                    *bool
}

func DefaultAdaptiveApplicationControlsListOperationOptions() AdaptiveApplicationControlsListOperationOptions {
	return AdaptiveApplicationControlsListOperationOptions{}
}

func (o AdaptiveApplicationControlsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o AdaptiveApplicationControlsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.IncludePathRecommendations != nil {
		out["includePathRecommendations"] = *o.IncludePathRecommendations
	}

	if o.Summary != nil {
		out["summary"] = *o.Summary
	}

	return out
}

// AdaptiveApplicationControlsList ...
func (c ApplicationWhitelistingsClient) AdaptiveApplicationControlsList(ctx context.Context, id commonids.SubscriptionId, options AdaptiveApplicationControlsListOperationOptions) (result AdaptiveApplicationControlsListOperationResponse, err error) {
	req, err := c.preparerForAdaptiveApplicationControlsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationwhitelistings.ApplicationWhitelistingsClient", "AdaptiveApplicationControlsList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationwhitelistings.ApplicationWhitelistingsClient", "AdaptiveApplicationControlsList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAdaptiveApplicationControlsList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationwhitelistings.ApplicationWhitelistingsClient", "AdaptiveApplicationControlsList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAdaptiveApplicationControlsList prepares the AdaptiveApplicationControlsList request.
func (c ApplicationWhitelistingsClient) preparerForAdaptiveApplicationControlsList(ctx context.Context, id commonids.SubscriptionId, options AdaptiveApplicationControlsListOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Security/applicationWhitelistings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAdaptiveApplicationControlsList handles the response to the AdaptiveApplicationControlsList request. The method always
// closes the http.Response Body.
func (c ApplicationWhitelistingsClient) responderForAdaptiveApplicationControlsList(resp *http.Response) (result AdaptiveApplicationControlsListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

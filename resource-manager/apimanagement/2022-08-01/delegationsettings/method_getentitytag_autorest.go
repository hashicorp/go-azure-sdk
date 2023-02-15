package delegationsettings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntityTagOperationResponse struct {
	HttpResponse *http.Response
}

// GetEntityTag ...
func (c DelegationSettingsClient) GetEntityTag(ctx context.Context, id ServiceId) (result GetEntityTagOperationResponse, err error) {
	req, err := c.preparerForGetEntityTag(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegationsettings.DelegationSettingsClient", "GetEntityTag", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegationsettings.DelegationSettingsClient", "GetEntityTag", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetEntityTag(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegationsettings.DelegationSettingsClient", "GetEntityTag", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetEntityTag prepares the GetEntityTag request.
func (c DelegationSettingsClient) preparerForGetEntityTag(ctx context.Context, id ServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsHead(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/portalsettings/delegation", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetEntityTag handles the response to the GetEntityTag request. The method always
// closes the http.Response Body.
func (c DelegationSettingsClient) responderForGetEntityTag(resp *http.Response) (result GetEntityTagOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

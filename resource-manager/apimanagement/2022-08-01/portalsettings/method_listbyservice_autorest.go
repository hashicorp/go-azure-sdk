package portalsettings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByServiceOperationResponse struct {
	HttpResponse *http.Response
	Model        *PortalSettingsCollection
}

// ListByService ...
func (c PortalSettingsClient) ListByService(ctx context.Context, id ServiceId) (result ListByServiceOperationResponse, err error) {
	req, err := c.preparerForListByService(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "portalsettings.PortalSettingsClient", "ListByService", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "portalsettings.PortalSettingsClient", "ListByService", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByService(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "portalsettings.PortalSettingsClient", "ListByService", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByService prepares the ListByService request.
func (c PortalSettingsClient) preparerForListByService(ctx context.Context, id ServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/portalsettings", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByService handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (c PortalSettingsClient) responderForListByService(resp *http.Response) (result ListByServiceOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

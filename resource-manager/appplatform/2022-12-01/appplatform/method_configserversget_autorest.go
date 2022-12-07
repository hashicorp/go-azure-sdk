package appplatform

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigServersGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConfigServerResource
}

// ConfigServersGet ...
func (c AppPlatformClient) ConfigServersGet(ctx context.Context, id SpringId) (result ConfigServersGetOperationResponse, err error) {
	req, err := c.preparerForConfigServersGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigServersGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigServersGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForConfigServersGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ConfigServersGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForConfigServersGet prepares the ConfigServersGet request.
func (c AppPlatformClient) preparerForConfigServersGet(ctx context.Context, id SpringId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/configServers/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForConfigServersGet handles the response to the ConfigServersGet request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForConfigServersGet(resp *http.Response) (result ConfigServersGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

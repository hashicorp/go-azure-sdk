package appliances

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

type GetTelemetryConfigOperationResponse struct {
	HttpResponse *http.Response
	Model        *ApplianceGetTelemetryConfigResult
}

// GetTelemetryConfig ...
func (c AppliancesClient) GetTelemetryConfig(ctx context.Context, id commonids.SubscriptionId) (result GetTelemetryConfigOperationResponse, err error) {
	req, err := c.preparerForGetTelemetryConfig(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appliances.AppliancesClient", "GetTelemetryConfig", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appliances.AppliancesClient", "GetTelemetryConfig", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetTelemetryConfig(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appliances.AppliancesClient", "GetTelemetryConfig", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetTelemetryConfig prepares the GetTelemetryConfig request.
func (c AppliancesClient) preparerForGetTelemetryConfig(ctx context.Context, id commonids.SubscriptionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.ResourceConnector/telemetryconfig", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetTelemetryConfig handles the response to the GetTelemetryConfig request. The method always
// closes the http.Response Body.
func (c AppliancesClient) responderForGetTelemetryConfig(resp *http.Response) (result GetTelemetryConfigOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package tenantconfigurationsyncstate

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantConfigurationGetSyncStateOperationResponse struct {
	HttpResponse *http.Response
	Model        *TenantConfigurationSyncStateContract
}

// TenantConfigurationGetSyncState ...
func (c TenantConfigurationSyncStateClient) TenantConfigurationGetSyncState(ctx context.Context, id ServiceId) (result TenantConfigurationGetSyncStateOperationResponse, err error) {
	req, err := c.preparerForTenantConfigurationGetSyncState(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantconfigurationsyncstate.TenantConfigurationSyncStateClient", "TenantConfigurationGetSyncState", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantconfigurationsyncstate.TenantConfigurationSyncStateClient", "TenantConfigurationGetSyncState", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTenantConfigurationGetSyncState(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantconfigurationsyncstate.TenantConfigurationSyncStateClient", "TenantConfigurationGetSyncState", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTenantConfigurationGetSyncState prepares the TenantConfigurationGetSyncState request.
func (c TenantConfigurationSyncStateClient) preparerForTenantConfigurationGetSyncState(ctx context.Context, id ServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/tenant/configuration/syncState", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForTenantConfigurationGetSyncState handles the response to the TenantConfigurationGetSyncState request. The method always
// closes the http.Response Body.
func (c TenantConfigurationSyncStateClient) responderForTenantConfigurationGetSyncState(resp *http.Response) (result TenantConfigurationGetSyncStateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

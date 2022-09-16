package sqlpoolsmaintenancewindows

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolMaintenanceWindowsCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
}

type SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions struct {
	MaintenanceWindowName *string
}

func DefaultSqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions() SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions {
	return SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions{}
}

func (o SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.MaintenanceWindowName != nil {
		out["maintenanceWindowName"] = *o.MaintenanceWindowName
	}

	return out
}

// SqlPoolMaintenanceWindowsCreateOrUpdate ...
func (c SqlPoolsMaintenanceWindowsClient) SqlPoolMaintenanceWindowsCreateOrUpdate(ctx context.Context, id SqlPoolId, input MaintenanceWindows, options SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions) (result SqlPoolMaintenanceWindowsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForSqlPoolMaintenanceWindowsCreateOrUpdate(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsmaintenancewindows.SqlPoolsMaintenanceWindowsClient", "SqlPoolMaintenanceWindowsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsmaintenancewindows.SqlPoolsMaintenanceWindowsClient", "SqlPoolMaintenanceWindowsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolMaintenanceWindowsCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsmaintenancewindows.SqlPoolsMaintenanceWindowsClient", "SqlPoolMaintenanceWindowsCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolMaintenanceWindowsCreateOrUpdate prepares the SqlPoolMaintenanceWindowsCreateOrUpdate request.
func (c SqlPoolsMaintenanceWindowsClient) preparerForSqlPoolMaintenanceWindowsCreateOrUpdate(ctx context.Context, id SqlPoolId, input MaintenanceWindows, options SqlPoolMaintenanceWindowsCreateOrUpdateOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/maintenancewindows/current", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolMaintenanceWindowsCreateOrUpdate handles the response to the SqlPoolMaintenanceWindowsCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c SqlPoolsMaintenanceWindowsClient) responderForSqlPoolMaintenanceWindowsCreateOrUpdate(resp *http.Response) (result SqlPoolMaintenanceWindowsCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

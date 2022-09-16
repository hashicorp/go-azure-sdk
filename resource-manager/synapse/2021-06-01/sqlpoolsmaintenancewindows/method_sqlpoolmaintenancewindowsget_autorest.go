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

type SqlPoolMaintenanceWindowsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *MaintenanceWindows
}

type SqlPoolMaintenanceWindowsGetOperationOptions struct {
	MaintenanceWindowName *string
}

func DefaultSqlPoolMaintenanceWindowsGetOperationOptions() SqlPoolMaintenanceWindowsGetOperationOptions {
	return SqlPoolMaintenanceWindowsGetOperationOptions{}
}

func (o SqlPoolMaintenanceWindowsGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o SqlPoolMaintenanceWindowsGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.MaintenanceWindowName != nil {
		out["maintenanceWindowName"] = *o.MaintenanceWindowName
	}

	return out
}

// SqlPoolMaintenanceWindowsGet ...
func (c SqlPoolsMaintenanceWindowsClient) SqlPoolMaintenanceWindowsGet(ctx context.Context, id SqlPoolId, options SqlPoolMaintenanceWindowsGetOperationOptions) (result SqlPoolMaintenanceWindowsGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolMaintenanceWindowsGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsmaintenancewindows.SqlPoolsMaintenanceWindowsClient", "SqlPoolMaintenanceWindowsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsmaintenancewindows.SqlPoolsMaintenanceWindowsClient", "SqlPoolMaintenanceWindowsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolMaintenanceWindowsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsmaintenancewindows.SqlPoolsMaintenanceWindowsClient", "SqlPoolMaintenanceWindowsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolMaintenanceWindowsGet prepares the SqlPoolMaintenanceWindowsGet request.
func (c SqlPoolsMaintenanceWindowsClient) preparerForSqlPoolMaintenanceWindowsGet(ctx context.Context, id SqlPoolId, options SqlPoolMaintenanceWindowsGetOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/maintenancewindows/current", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolMaintenanceWindowsGet handles the response to the SqlPoolMaintenanceWindowsGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsMaintenanceWindowsClient) responderForSqlPoolMaintenanceWindowsGet(resp *http.Response) (result SqlPoolMaintenanceWindowsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

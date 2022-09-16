package sqlpoolsmaintenancewindowoptions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolMaintenanceWindowOptionsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *MaintenanceWindowOptions
}

type SqlPoolMaintenanceWindowOptionsGetOperationOptions struct {
	MaintenanceWindowOptionsName *string
}

func DefaultSqlPoolMaintenanceWindowOptionsGetOperationOptions() SqlPoolMaintenanceWindowOptionsGetOperationOptions {
	return SqlPoolMaintenanceWindowOptionsGetOperationOptions{}
}

func (o SqlPoolMaintenanceWindowOptionsGetOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o SqlPoolMaintenanceWindowOptionsGetOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.MaintenanceWindowOptionsName != nil {
		out["maintenanceWindowOptionsName"] = *o.MaintenanceWindowOptionsName
	}

	return out
}

// SqlPoolMaintenanceWindowOptionsGet ...
func (c SqlPoolsMaintenanceWindowOptionsClient) SqlPoolMaintenanceWindowOptionsGet(ctx context.Context, id SqlPoolId, options SqlPoolMaintenanceWindowOptionsGetOperationOptions) (result SqlPoolMaintenanceWindowOptionsGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolMaintenanceWindowOptionsGet(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsmaintenancewindowoptions.SqlPoolsMaintenanceWindowOptionsClient", "SqlPoolMaintenanceWindowOptionsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsmaintenancewindowoptions.SqlPoolsMaintenanceWindowOptionsClient", "SqlPoolMaintenanceWindowOptionsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolMaintenanceWindowOptionsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsmaintenancewindowoptions.SqlPoolsMaintenanceWindowOptionsClient", "SqlPoolMaintenanceWindowOptionsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolMaintenanceWindowOptionsGet prepares the SqlPoolMaintenanceWindowOptionsGet request.
func (c SqlPoolsMaintenanceWindowOptionsClient) preparerForSqlPoolMaintenanceWindowOptionsGet(ctx context.Context, id SqlPoolId, options SqlPoolMaintenanceWindowOptionsGetOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/maintenanceWindowOptions/current", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolMaintenanceWindowOptionsGet handles the response to the SqlPoolMaintenanceWindowOptionsGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsMaintenanceWindowOptionsClient) responderForSqlPoolMaintenanceWindowOptionsGet(resp *http.Response) (result SqlPoolMaintenanceWindowOptionsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

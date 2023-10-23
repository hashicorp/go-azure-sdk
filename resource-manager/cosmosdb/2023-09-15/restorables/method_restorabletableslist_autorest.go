package restorables

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableTablesListOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorableTablesListResult
}

type RestorableTablesListOperationOptions struct {
	EndTime   *string
	StartTime *string
}

func DefaultRestorableTablesListOperationOptions() RestorableTablesListOperationOptions {
	return RestorableTablesListOperationOptions{}
}

func (o RestorableTablesListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RestorableTablesListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.EndTime != nil {
		out["endTime"] = *o.EndTime
	}

	if o.StartTime != nil {
		out["startTime"] = *o.StartTime
	}

	return out
}

// RestorableTablesList ...
func (c RestorablesClient) RestorableTablesList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableTablesListOperationOptions) (result RestorableTablesListOperationResponse, err error) {
	req, err := c.preparerForRestorableTablesList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableTablesList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableTablesList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestorableTablesList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableTablesList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestorableTablesList prepares the RestorableTablesList request.
func (c RestorablesClient) preparerForRestorableTablesList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableTablesListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/restorableTables", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestorableTablesList handles the response to the RestorableTablesList request. The method always
// closes the http.Response Body.
func (c RestorablesClient) responderForRestorableTablesList(resp *http.Response) (result RestorableTablesListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

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

type RestorableSqlContainersListOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorableSqlContainersListResult
}

type RestorableSqlContainersListOperationOptions struct {
	EndTime                  *string
	RestorableSqlDatabaseRid *string
	StartTime                *string
}

func DefaultRestorableSqlContainersListOperationOptions() RestorableSqlContainersListOperationOptions {
	return RestorableSqlContainersListOperationOptions{}
}

func (o RestorableSqlContainersListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RestorableSqlContainersListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.EndTime != nil {
		out["endTime"] = *o.EndTime
	}

	if o.RestorableSqlDatabaseRid != nil {
		out["restorableSqlDatabaseRid"] = *o.RestorableSqlDatabaseRid
	}

	if o.StartTime != nil {
		out["startTime"] = *o.StartTime
	}

	return out
}

// RestorableSqlContainersList ...
func (c RestorablesClient) RestorableSqlContainersList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableSqlContainersListOperationOptions) (result RestorableSqlContainersListOperationResponse, err error) {
	req, err := c.preparerForRestorableSqlContainersList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableSqlContainersList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableSqlContainersList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestorableSqlContainersList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableSqlContainersList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestorableSqlContainersList prepares the RestorableSqlContainersList request.
func (c RestorablesClient) preparerForRestorableSqlContainersList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableSqlContainersListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/restorableSqlContainers", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestorableSqlContainersList handles the response to the RestorableSqlContainersList request. The method always
// closes the http.Response Body.
func (c RestorablesClient) responderForRestorableSqlContainersList(resp *http.Response) (result RestorableSqlContainersListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

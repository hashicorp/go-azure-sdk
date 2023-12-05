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

type RestorableMongodbCollectionsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *RestorableMongodbCollectionsListResult
}

type RestorableMongodbCollectionsListOperationOptions struct {
	EndTime                      *string
	RestorableMongodbDatabaseRid *string
	StartTime                    *string
}

func DefaultRestorableMongodbCollectionsListOperationOptions() RestorableMongodbCollectionsListOperationOptions {
	return RestorableMongodbCollectionsListOperationOptions{}
}

func (o RestorableMongodbCollectionsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o RestorableMongodbCollectionsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.EndTime != nil {
		out["endTime"] = *o.EndTime
	}

	if o.RestorableMongodbDatabaseRid != nil {
		out["restorableMongodbDatabaseRid"] = *o.RestorableMongodbDatabaseRid
	}

	if o.StartTime != nil {
		out["startTime"] = *o.StartTime
	}

	return out
}

// RestorableMongodbCollectionsList ...
func (c RestorablesClient) RestorableMongodbCollectionsList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableMongodbCollectionsListOperationOptions) (result RestorableMongodbCollectionsListOperationResponse, err error) {
	req, err := c.preparerForRestorableMongodbCollectionsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableMongodbCollectionsList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableMongodbCollectionsList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRestorableMongodbCollectionsList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restorables.RestorablesClient", "RestorableMongodbCollectionsList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRestorableMongodbCollectionsList prepares the RestorableMongodbCollectionsList request.
func (c RestorablesClient) preparerForRestorableMongodbCollectionsList(ctx context.Context, id RestorableDatabaseAccountId, options RestorableMongodbCollectionsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/restorableMongodbCollections", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRestorableMongodbCollectionsList handles the response to the RestorableMongodbCollectionsList request. The method always
// closes the http.Response Body.
func (c RestorablesClient) responderForRestorableMongodbCollectionsList(resp *http.Response) (result RestorableMongodbCollectionsListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

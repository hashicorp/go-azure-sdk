package bms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobsExportOperationResponse struct {
	HttpResponse *http.Response
}

type JobsExportOperationOptions struct {
	Filter *string
}

func DefaultJobsExportOperationOptions() JobsExportOperationOptions {
	return JobsExportOperationOptions{}
}

func (o JobsExportOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o JobsExportOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// JobsExport ...
func (c BmsClient) JobsExport(ctx context.Context, id VaultId, options JobsExportOperationOptions) (result JobsExportOperationResponse, err error) {
	req, err := c.preparerForJobsExport(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "JobsExport", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "JobsExport", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForJobsExport(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "JobsExport", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForJobsExport prepares the JobsExport request.
func (c BmsClient) preparerForJobsExport(ctx context.Context, id VaultId, options JobsExportOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/backupJobsExport", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForJobsExport handles the response to the JobsExport request. The method always
// closes the http.Response Body.
func (c BmsClient) responderForJobsExport(resp *http.Response) (result JobsExportOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

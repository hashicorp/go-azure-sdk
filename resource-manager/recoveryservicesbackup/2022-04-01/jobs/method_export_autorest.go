package jobs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportOperationResponse struct {
	HttpResponse *http.Response
}

type ExportOperationOptions struct {
	Filter *string
}

func DefaultExportOperationOptions() ExportOperationOptions {
	return ExportOperationOptions{}
}

func (o ExportOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ExportOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// Export ...
func (c JobsClient) Export(ctx context.Context, id VaultId, options ExportOperationOptions) (result ExportOperationResponse, err error) {
	req, err := c.preparerForExport(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobs.JobsClient", "Export", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobs.JobsClient", "Export", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForExport(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobs.JobsClient", "Export", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForExport prepares the Export request.
func (c JobsClient) preparerForExport(ctx context.Context, id VaultId, options ExportOperationOptions) (*http.Request, error) {
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

// responderForExport handles the response to the Export request. The method always
// closes the http.Response Body.
func (c JobsClient) responderForExport(resp *http.Response) (result ExportOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

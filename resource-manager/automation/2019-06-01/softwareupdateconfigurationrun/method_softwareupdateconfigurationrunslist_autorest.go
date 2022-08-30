package softwareupdateconfigurationrun

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftwareUpdateConfigurationRunsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *SoftwareUpdateConfigurationRunListResult
}

type SoftwareUpdateConfigurationRunsListOperationOptions struct {
	ClientRequestId *string
	Filter          *string
	Skip            *string
	Top             *string
}

func DefaultSoftwareUpdateConfigurationRunsListOperationOptions() SoftwareUpdateConfigurationRunsListOperationOptions {
	return SoftwareUpdateConfigurationRunsListOperationOptions{}
}

func (o SoftwareUpdateConfigurationRunsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ClientRequestId != nil {
		out["clientRequestId"] = *o.ClientRequestId
	}

	return out
}

func (o SoftwareUpdateConfigurationRunsListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Skip != nil {
		out["$skip"] = *o.Skip
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// SoftwareUpdateConfigurationRunsList ...
func (c SoftwareUpdateConfigurationRunClient) SoftwareUpdateConfigurationRunsList(ctx context.Context, id AutomationAccountId, options SoftwareUpdateConfigurationRunsListOperationOptions) (result SoftwareUpdateConfigurationRunsListOperationResponse, err error) {
	req, err := c.preparerForSoftwareUpdateConfigurationRunsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationrun.SoftwareUpdateConfigurationRunClient", "SoftwareUpdateConfigurationRunsList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationrun.SoftwareUpdateConfigurationRunClient", "SoftwareUpdateConfigurationRunsList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSoftwareUpdateConfigurationRunsList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationrun.SoftwareUpdateConfigurationRunClient", "SoftwareUpdateConfigurationRunsList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSoftwareUpdateConfigurationRunsList prepares the SoftwareUpdateConfigurationRunsList request.
func (c SoftwareUpdateConfigurationRunClient) preparerForSoftwareUpdateConfigurationRunsList(ctx context.Context, id AutomationAccountId, options SoftwareUpdateConfigurationRunsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/softwareUpdateConfigurationRuns", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSoftwareUpdateConfigurationRunsList handles the response to the SoftwareUpdateConfigurationRunsList request. The method always
// closes the http.Response Body.
func (c SoftwareUpdateConfigurationRunClient) responderForSoftwareUpdateConfigurationRunsList(resp *http.Response) (result SoftwareUpdateConfigurationRunsListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

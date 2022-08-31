package softwareupdateconfigurationmachinerun

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftwareUpdateConfigurationMachineRunsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *SoftwareUpdateConfigurationMachineRunListResult
}

type SoftwareUpdateConfigurationMachineRunsListOperationOptions struct {
	ClientRequestId *string
	Filter          *string
	Skip            *string
	Top             *string
}

func DefaultSoftwareUpdateConfigurationMachineRunsListOperationOptions() SoftwareUpdateConfigurationMachineRunsListOperationOptions {
	return SoftwareUpdateConfigurationMachineRunsListOperationOptions{}
}

func (o SoftwareUpdateConfigurationMachineRunsListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ClientRequestId != nil {
		out["clientRequestId"] = *o.ClientRequestId
	}

	return out
}

func (o SoftwareUpdateConfigurationMachineRunsListOperationOptions) toQueryString() map[string]interface{} {
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

// SoftwareUpdateConfigurationMachineRunsList ...
func (c SoftwareUpdateConfigurationMachineRunClient) SoftwareUpdateConfigurationMachineRunsList(ctx context.Context, id AutomationAccountId, options SoftwareUpdateConfigurationMachineRunsListOperationOptions) (result SoftwareUpdateConfigurationMachineRunsListOperationResponse, err error) {
	req, err := c.preparerForSoftwareUpdateConfigurationMachineRunsList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationmachinerun.SoftwareUpdateConfigurationMachineRunClient", "SoftwareUpdateConfigurationMachineRunsList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationmachinerun.SoftwareUpdateConfigurationMachineRunClient", "SoftwareUpdateConfigurationMachineRunsList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSoftwareUpdateConfigurationMachineRunsList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationmachinerun.SoftwareUpdateConfigurationMachineRunClient", "SoftwareUpdateConfigurationMachineRunsList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSoftwareUpdateConfigurationMachineRunsList prepares the SoftwareUpdateConfigurationMachineRunsList request.
func (c SoftwareUpdateConfigurationMachineRunClient) preparerForSoftwareUpdateConfigurationMachineRunsList(ctx context.Context, id AutomationAccountId, options SoftwareUpdateConfigurationMachineRunsListOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/softwareUpdateConfigurationMachineRuns", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSoftwareUpdateConfigurationMachineRunsList handles the response to the SoftwareUpdateConfigurationMachineRunsList request. The method always
// closes the http.Response Body.
func (c SoftwareUpdateConfigurationMachineRunClient) responderForSoftwareUpdateConfigurationMachineRunsList(resp *http.Response) (result SoftwareUpdateConfigurationMachineRunsListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package softwareupdateconfigurationmachinerun

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftwareUpdateConfigurationMachineRunsGetByIdOperationResponse struct {
	HttpResponse *http.Response
	Model        *SoftwareUpdateConfigurationMachineRun
}

type SoftwareUpdateConfigurationMachineRunsGetByIdOperationOptions struct {
	ClientRequestId *string
}

func DefaultSoftwareUpdateConfigurationMachineRunsGetByIdOperationOptions() SoftwareUpdateConfigurationMachineRunsGetByIdOperationOptions {
	return SoftwareUpdateConfigurationMachineRunsGetByIdOperationOptions{}
}

func (o SoftwareUpdateConfigurationMachineRunsGetByIdOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ClientRequestId != nil {
		out["clientRequestId"] = *o.ClientRequestId
	}

	return out
}

func (o SoftwareUpdateConfigurationMachineRunsGetByIdOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// SoftwareUpdateConfigurationMachineRunsGetById ...
func (c SoftwareUpdateConfigurationMachineRunClient) SoftwareUpdateConfigurationMachineRunsGetById(ctx context.Context, id SoftwareUpdateConfigurationMachineRunId, options SoftwareUpdateConfigurationMachineRunsGetByIdOperationOptions) (result SoftwareUpdateConfigurationMachineRunsGetByIdOperationResponse, err error) {
	req, err := c.preparerForSoftwareUpdateConfigurationMachineRunsGetById(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationmachinerun.SoftwareUpdateConfigurationMachineRunClient", "SoftwareUpdateConfigurationMachineRunsGetById", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationmachinerun.SoftwareUpdateConfigurationMachineRunClient", "SoftwareUpdateConfigurationMachineRunsGetById", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSoftwareUpdateConfigurationMachineRunsGetById(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationmachinerun.SoftwareUpdateConfigurationMachineRunClient", "SoftwareUpdateConfigurationMachineRunsGetById", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSoftwareUpdateConfigurationMachineRunsGetById prepares the SoftwareUpdateConfigurationMachineRunsGetById request.
func (c SoftwareUpdateConfigurationMachineRunClient) preparerForSoftwareUpdateConfigurationMachineRunsGetById(ctx context.Context, id SoftwareUpdateConfigurationMachineRunId, options SoftwareUpdateConfigurationMachineRunsGetByIdOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSoftwareUpdateConfigurationMachineRunsGetById handles the response to the SoftwareUpdateConfigurationMachineRunsGetById request. The method always
// closes the http.Response Body.
func (c SoftwareUpdateConfigurationMachineRunClient) responderForSoftwareUpdateConfigurationMachineRunsGetById(resp *http.Response) (result SoftwareUpdateConfigurationMachineRunsGetByIdOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

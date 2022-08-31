package softwareupdateconfigurationrun

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftwareUpdateConfigurationRunsGetByIdOperationResponse struct {
	HttpResponse *http.Response
	Model        *SoftwareUpdateConfigurationRun
}

type SoftwareUpdateConfigurationRunsGetByIdOperationOptions struct {
	ClientRequestId *string
}

func DefaultSoftwareUpdateConfigurationRunsGetByIdOperationOptions() SoftwareUpdateConfigurationRunsGetByIdOperationOptions {
	return SoftwareUpdateConfigurationRunsGetByIdOperationOptions{}
}

func (o SoftwareUpdateConfigurationRunsGetByIdOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ClientRequestId != nil {
		out["clientRequestId"] = *o.ClientRequestId
	}

	return out
}

func (o SoftwareUpdateConfigurationRunsGetByIdOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// SoftwareUpdateConfigurationRunsGetById ...
func (c SoftwareUpdateConfigurationRunClient) SoftwareUpdateConfigurationRunsGetById(ctx context.Context, id SoftwareUpdateConfigurationRunId, options SoftwareUpdateConfigurationRunsGetByIdOperationOptions) (result SoftwareUpdateConfigurationRunsGetByIdOperationResponse, err error) {
	req, err := c.preparerForSoftwareUpdateConfigurationRunsGetById(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationrun.SoftwareUpdateConfigurationRunClient", "SoftwareUpdateConfigurationRunsGetById", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationrun.SoftwareUpdateConfigurationRunClient", "SoftwareUpdateConfigurationRunsGetById", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSoftwareUpdateConfigurationRunsGetById(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationrun.SoftwareUpdateConfigurationRunClient", "SoftwareUpdateConfigurationRunsGetById", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSoftwareUpdateConfigurationRunsGetById prepares the SoftwareUpdateConfigurationRunsGetById request.
func (c SoftwareUpdateConfigurationRunClient) preparerForSoftwareUpdateConfigurationRunsGetById(ctx context.Context, id SoftwareUpdateConfigurationRunId, options SoftwareUpdateConfigurationRunsGetByIdOperationOptions) (*http.Request, error) {
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

// responderForSoftwareUpdateConfigurationRunsGetById handles the response to the SoftwareUpdateConfigurationRunsGetById request. The method always
// closes the http.Response Body.
func (c SoftwareUpdateConfigurationRunClient) responderForSoftwareUpdateConfigurationRunsGetById(resp *http.Response) (result SoftwareUpdateConfigurationRunsGetByIdOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

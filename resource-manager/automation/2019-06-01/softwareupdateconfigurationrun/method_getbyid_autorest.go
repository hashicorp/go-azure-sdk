package softwareupdateconfigurationrun

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetByIdOperationResponse struct {
	HttpResponse *http.Response
	Model        *SoftwareUpdateConfigurationRun
}

type GetByIdOperationOptions struct {
	ClientRequestId *string
}

func DefaultGetByIdOperationOptions() GetByIdOperationOptions {
	return GetByIdOperationOptions{}
}

func (o GetByIdOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.ClientRequestId != nil {
		out["clientRequestId"] = *o.ClientRequestId
	}

	return out
}

func (o GetByIdOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// GetById ...
func (c SoftwareUpdateConfigurationRunClient) GetById(ctx context.Context, id SoftwareUpdateConfigurationRunId, options GetByIdOperationOptions) (result GetByIdOperationResponse, err error) {
	req, err := c.preparerForGetById(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationrun.SoftwareUpdateConfigurationRunClient", "GetById", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationrun.SoftwareUpdateConfigurationRunClient", "GetById", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetById(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "softwareupdateconfigurationrun.SoftwareUpdateConfigurationRunClient", "GetById", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetById prepares the GetById request.
func (c SoftwareUpdateConfigurationRunClient) preparerForGetById(ctx context.Context, id SoftwareUpdateConfigurationRunId, options GetByIdOperationOptions) (*http.Request, error) {
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

// responderForGetById handles the response to the GetById request. The method always
// closes the http.Response Body.
func (c SoftwareUpdateConfigurationRunClient) responderForGetById(resp *http.Response) (result GetByIdOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package dataconnectors

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorsCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *DataConnector
}

// DataConnectorsCreateOrUpdate ...
func (c DataConnectorsClient) DataConnectorsCreateOrUpdate(ctx context.Context, id DataConnectorId, input DataConnector) (result DataConnectorsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForDataConnectorsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataconnectors.DataConnectorsClient", "DataConnectorsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataconnectors.DataConnectorsClient", "DataConnectorsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDataConnectorsCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataconnectors.DataConnectorsClient", "DataConnectorsCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDataConnectorsCreateOrUpdate prepares the DataConnectorsCreateOrUpdate request.
func (c DataConnectorsClient) preparerForDataConnectorsCreateOrUpdate(ctx context.Context, id DataConnectorId, input DataConnector) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDataConnectorsCreateOrUpdate handles the response to the DataConnectorsCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c DataConnectorsClient) responderForDataConnectorsCreateOrUpdate(resp *http.Response) (result DataConnectorsCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, fmt.Errorf("reading response body for DataConnector: %+v", err)
	}
	model, err := unmarshalDataConnectorImplementation(b)
	if err != nil {
		return
	}
	result.Model = &model
	return
}

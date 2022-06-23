package checkdataconnectorrequirements

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorsCheckRequirementsPostOperationResponse struct {
	HttpResponse *http.Response
	Model        *DataConnectorRequirementsState
}

// DataConnectorsCheckRequirementsPost ...
func (c CheckDataConnectorRequirementsClient) DataConnectorsCheckRequirementsPost(ctx context.Context, id WorkspaceId, input DataConnectorsCheckRequirements) (result DataConnectorsCheckRequirementsPostOperationResponse, err error) {
	req, err := c.preparerForDataConnectorsCheckRequirementsPost(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "checkdataconnectorrequirements.CheckDataConnectorRequirementsClient", "DataConnectorsCheckRequirementsPost", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "checkdataconnectorrequirements.CheckDataConnectorRequirementsClient", "DataConnectorsCheckRequirementsPost", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDataConnectorsCheckRequirementsPost(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "checkdataconnectorrequirements.CheckDataConnectorRequirementsClient", "DataConnectorsCheckRequirementsPost", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDataConnectorsCheckRequirementsPost prepares the DataConnectorsCheckRequirementsPost request.
func (c CheckDataConnectorRequirementsClient) preparerForDataConnectorsCheckRequirementsPost(ctx context.Context, id WorkspaceId, input DataConnectorsCheckRequirements) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/dataConnectorsCheckRequirements", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDataConnectorsCheckRequirementsPost handles the response to the DataConnectorsCheckRequirementsPost request. The method always
// closes the http.Response Body.
func (c CheckDataConnectorRequirementsClient) responderForDataConnectorsCheckRequirementsPost(resp *http.Response) (result DataConnectorsCheckRequirementsPostOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

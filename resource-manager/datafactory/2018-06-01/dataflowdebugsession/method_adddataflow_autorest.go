package dataflowdebugsession

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddDataFlowOperationResponse struct {
	HttpResponse *http.Response
	Model        *AddDataFlowToDebugSessionResponse
}

// AddDataFlow ...
func (c DataFlowDebugSessionClient) AddDataFlow(ctx context.Context, id FactoryId, input DataFlowDebugPackage) (result AddDataFlowOperationResponse, err error) {
	req, err := c.preparerForAddDataFlow(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataflowdebugsession.DataFlowDebugSessionClient", "AddDataFlow", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataflowdebugsession.DataFlowDebugSessionClient", "AddDataFlow", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAddDataFlow(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataflowdebugsession.DataFlowDebugSessionClient", "AddDataFlow", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAddDataFlow prepares the AddDataFlow request.
func (c DataFlowDebugSessionClient) preparerForAddDataFlow(ctx context.Context, id FactoryId, input DataFlowDebugPackage) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/addDataFlowToDebugSession", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAddDataFlow handles the response to the AddDataFlow request. The method always
// closes the http.Response Body.
func (c DataFlowDebugSessionClient) responderForAddDataFlow(resp *http.Response) (result AddDataFlowOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

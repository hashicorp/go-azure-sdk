package servicetaskresource

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceTasksUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ProjectTask
}

// ServiceTasksUpdate ...
func (c ServiceTaskResourceClient) ServiceTasksUpdate(ctx context.Context, id ServiceTaskId, input ProjectTask) (result ServiceTasksUpdateOperationResponse, err error) {
	req, err := c.preparerForServiceTasksUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicetaskresource.ServiceTaskResourceClient", "ServiceTasksUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicetaskresource.ServiceTaskResourceClient", "ServiceTasksUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForServiceTasksUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicetaskresource.ServiceTaskResourceClient", "ServiceTasksUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForServiceTasksUpdate prepares the ServiceTasksUpdate request.
func (c ServiceTaskResourceClient) preparerForServiceTasksUpdate(ctx context.Context, id ServiceTaskId, input ProjectTask) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForServiceTasksUpdate handles the response to the ServiceTasksUpdate request. The method always
// closes the http.Response Body.
func (c ServiceTaskResourceClient) responderForServiceTasksUpdate(resp *http.Response) (result ServiceTasksUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

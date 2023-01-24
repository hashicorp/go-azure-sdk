package standardoperation

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *Project
}

// ProjectsGet ...
func (c StandardOperationClient) ProjectsGet(ctx context.Context, id ProjectId) (result ProjectsGetOperationResponse, err error) {
	req, err := c.preparerForProjectsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ProjectsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ProjectsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProjectsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "standardoperation.StandardOperationClient", "ProjectsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProjectsGet prepares the ProjectsGet request.
func (c StandardOperationClient) preparerForProjectsGet(ctx context.Context, id ProjectId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForProjectsGet handles the response to the ProjectsGet request. The method always
// closes the http.Response Body.
func (c StandardOperationClient) responderForProjectsGet(resp *http.Response) (result ProjectsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

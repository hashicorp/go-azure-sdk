package jobs

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetJobOperationResponse struct {
	HttpResponse *http.Response
	Model        *VMwareJob
}

// GetJob ...
func (c JobsClient) GetJob(ctx context.Context, id commonids.VMwareSiteJobId) (result GetJobOperationResponse, err error) {
	req, err := c.preparerForGetJob(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobs.JobsClient", "GetJob", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobs.JobsClient", "GetJob", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetJob(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "jobs.JobsClient", "GetJob", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetJob prepares the GetJob request.
func (c JobsClient) preparerForGetJob(ctx context.Context, id commonids.VMwareSiteJobId) (*http.Request, error) {
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

// responderForGetJob handles the response to the GetJob request. The method always
// closes the http.Response Body.
func (c JobsClient) responderForGetJob(resp *http.Response) (result GetJobOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

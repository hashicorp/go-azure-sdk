package onboardtod4api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type APICollectionsOnboardAzureApiManagementApiOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
	Model        *ApiCollection
}

// APICollectionsOnboardAzureApiManagementApi ...
func (c OnboardToD4APIClient) APICollectionsOnboardAzureApiManagementApi(ctx context.Context, id ApiCollectionId) (result APICollectionsOnboardAzureApiManagementApiOperationResponse, err error) {
	req, err := c.preparerForAPICollectionsOnboardAzureApiManagementApi(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "onboardtod4api.OnboardToD4APIClient", "APICollectionsOnboardAzureApiManagementApi", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForAPICollectionsOnboardAzureApiManagementApi(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "onboardtod4api.OnboardToD4APIClient", "APICollectionsOnboardAzureApiManagementApi", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// APICollectionsOnboardAzureApiManagementApiThenPoll performs APICollectionsOnboardAzureApiManagementApi then polls until it's completed
func (c OnboardToD4APIClient) APICollectionsOnboardAzureApiManagementApiThenPoll(ctx context.Context, id ApiCollectionId) error {
	result, err := c.APICollectionsOnboardAzureApiManagementApi(ctx, id)
	if err != nil {
		return fmt.Errorf("performing APICollectionsOnboardAzureApiManagementApi: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after APICollectionsOnboardAzureApiManagementApi: %+v", err)
	}

	return nil
}

// preparerForAPICollectionsOnboardAzureApiManagementApi prepares the APICollectionsOnboardAzureApiManagementApi request.
func (c OnboardToD4APIClient) preparerForAPICollectionsOnboardAzureApiManagementApi(ctx context.Context, id ApiCollectionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForAPICollectionsOnboardAzureApiManagementApi sends the APICollectionsOnboardAzureApiManagementApi request. The method will close the
// http.Response Body if it receives an error.
func (c OnboardToD4APIClient) senderForAPICollectionsOnboardAzureApiManagementApi(ctx context.Context, req *http.Request) (future APICollectionsOnboardAzureApiManagementApiOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}

package appserviceenvironments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TestUpgradeAvailableNotificationOperationResponse struct {
	HttpResponse *http.Response
}

// TestUpgradeAvailableNotification ...
func (c AppServiceEnvironmentsClient) TestUpgradeAvailableNotification(ctx context.Context, id HostingEnvironmentId) (result TestUpgradeAvailableNotificationOperationResponse, err error) {
	req, err := c.preparerForTestUpgradeAvailableNotification(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "TestUpgradeAvailableNotification", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "TestUpgradeAvailableNotification", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTestUpgradeAvailableNotification(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "TestUpgradeAvailableNotification", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTestUpgradeAvailableNotification prepares the TestUpgradeAvailableNotification request.
func (c AppServiceEnvironmentsClient) preparerForTestUpgradeAvailableNotification(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/testUpgradeAvailableNotification", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForTestUpgradeAvailableNotification handles the response to the TestUpgradeAvailableNotification request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForTestUpgradeAvailableNotification(resp *http.Response) (result TestUpgradeAvailableNotificationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

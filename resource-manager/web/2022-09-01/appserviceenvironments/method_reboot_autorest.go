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

type RebootOperationResponse struct {
	HttpResponse *http.Response
}

// Reboot ...
func (c AppServiceEnvironmentsClient) Reboot(ctx context.Context, id HostingEnvironmentId) (result RebootOperationResponse, err error) {
	req, err := c.preparerForReboot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "Reboot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "Reboot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForReboot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "Reboot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForReboot prepares the Reboot request.
func (c AppServiceEnvironmentsClient) preparerForReboot(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/reboot", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForReboot handles the response to the Reboot request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForReboot(resp *http.Response) (result RebootOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

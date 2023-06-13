package recoverypointsgetaccesstoken

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointsGetAccessTokenOperationResponse struct {
	HttpResponse *http.Response
	Model        *CrrAccessTokenResource
}

// RecoveryPointsGetAccessToken ...
func (c RecoveryPointsGetAccessTokenClient) RecoveryPointsGetAccessToken(ctx context.Context, id RecoveryPointId, input AADPropertiesResource) (result RecoveryPointsGetAccessTokenOperationResponse, err error) {
	req, err := c.preparerForRecoveryPointsGetAccessToken(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverypointsgetaccesstoken.RecoveryPointsGetAccessTokenClient", "RecoveryPointsGetAccessToken", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverypointsgetaccesstoken.RecoveryPointsGetAccessTokenClient", "RecoveryPointsGetAccessToken", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRecoveryPointsGetAccessToken(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoverypointsgetaccesstoken.RecoveryPointsGetAccessTokenClient", "RecoveryPointsGetAccessToken", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRecoveryPointsGetAccessToken prepares the RecoveryPointsGetAccessToken request.
func (c RecoveryPointsGetAccessTokenClient) preparerForRecoveryPointsGetAccessToken(ctx context.Context, id RecoveryPointId, input AADPropertiesResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/accessToken", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRecoveryPointsGetAccessToken handles the response to the RecoveryPointsGetAccessToken request. The method always
// closes the http.Response Body.
func (c RecoveryPointsGetAccessTokenClient) responderForRecoveryPointsGetAccessToken(resp *http.Response) (result RecoveryPointsGetAccessTokenOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

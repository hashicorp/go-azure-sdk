package post

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotDpsResourceListKeysForKeyNameOperationResponse struct {
	HttpResponse *http.Response
	Model        *SharedAccessSignatureAuthorizationRuleAccessRightsDescription
}

// IotDpsResourceListKeysForKeyName ...
func (c POSTClient) IotDpsResourceListKeysForKeyName(ctx context.Context, id KeyId) (result IotDpsResourceListKeysForKeyNameOperationResponse, err error) {
	req, err := c.preparerForIotDpsResourceListKeysForKeyName(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "IotDpsResourceListKeysForKeyName", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "IotDpsResourceListKeysForKeyName", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIotDpsResourceListKeysForKeyName(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "IotDpsResourceListKeysForKeyName", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIotDpsResourceListKeysForKeyName prepares the IotDpsResourceListKeysForKeyName request.
func (c POSTClient) preparerForIotDpsResourceListKeysForKeyName(ctx context.Context, id KeyId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listkeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIotDpsResourceListKeysForKeyName handles the response to the IotDpsResourceListKeysForKeyName request. The method always
// closes the http.Response Body.
func (c POSTClient) responderForIotDpsResourceListKeysForKeyName(resp *http.Response) (result IotDpsResourceListKeysForKeyNameOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

package delegationsettings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSecretsOperationResponse struct {
	HttpResponse *http.Response
	Model        *PortalSettingValidationKeyContract
}

// ListSecrets ...
func (c DelegationSettingsClient) ListSecrets(ctx context.Context, id ServiceId) (result ListSecretsOperationResponse, err error) {
	req, err := c.preparerForListSecrets(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegationsettings.DelegationSettingsClient", "ListSecrets", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegationsettings.DelegationSettingsClient", "ListSecrets", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListSecrets(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delegationsettings.DelegationSettingsClient", "ListSecrets", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListSecrets prepares the ListSecrets request.
func (c DelegationSettingsClient) preparerForListSecrets(ctx context.Context, id ServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/portalsettings/delegation/listSecrets", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListSecrets handles the response to the ListSecrets request. The method always
// closes the http.Response Body.
func (c DelegationSettingsClient) responderForListSecrets(resp *http.Response) (result ListSecretsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

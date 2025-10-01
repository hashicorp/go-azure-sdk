package bms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionIntentValidateOperationResponse struct {
	HttpResponse *http.Response
	Model        *PreValidateEnableBackupResponse
}

// ProtectionIntentValidate ...
func (c BmsClient) ProtectionIntentValidate(ctx context.Context, id LocationId, input PreValidateEnableBackupRequest) (result ProtectionIntentValidateOperationResponse, err error) {
	req, err := c.preparerForProtectionIntentValidate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ProtectionIntentValidate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ProtectionIntentValidate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForProtectionIntentValidate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ProtectionIntentValidate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForProtectionIntentValidate prepares the ProtectionIntentValidate request.
func (c BmsClient) preparerForProtectionIntentValidate(ctx context.Context, id LocationId, input PreValidateEnableBackupRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backupPreValidateProtection", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForProtectionIntentValidate handles the response to the ProtectionIntentValidate request. The method always
// closes the http.Response Body.
func (c BmsClient) responderForProtectionIntentValidate(resp *http.Response) (result ProtectionIntentValidateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

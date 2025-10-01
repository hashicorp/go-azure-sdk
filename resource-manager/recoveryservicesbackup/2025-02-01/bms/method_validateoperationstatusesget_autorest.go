package bms

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateOperationStatusesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *OperationStatus
}

// ValidateOperationStatusesGet ...
func (c BmsClient) ValidateOperationStatusesGet(ctx context.Context, id BackupValidateOperationsStatusId) (result ValidateOperationStatusesGetOperationResponse, err error) {
	req, err := c.preparerForValidateOperationStatusesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ValidateOperationStatusesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ValidateOperationStatusesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForValidateOperationStatusesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "ValidateOperationStatusesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForValidateOperationStatusesGet prepares the ValidateOperationStatusesGet request.
func (c BmsClient) preparerForValidateOperationStatusesGet(ctx context.Context, id BackupValidateOperationsStatusId) (*http.Request, error) {
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

// responderForValidateOperationStatusesGet handles the response to the ValidateOperationStatusesGet request. The method always
// closes the http.Response Body.
func (c BmsClient) responderForValidateOperationStatusesGet(resp *http.Response) (result ValidateOperationStatusesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

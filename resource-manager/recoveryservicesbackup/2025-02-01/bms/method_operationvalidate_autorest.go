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

type OperationValidateOperationResponse struct {
	HttpResponse *http.Response
	Model        *ValidateOperationsResponse
}

// OperationValidate ...
func (c BmsClient) OperationValidate(ctx context.Context, id VaultId, input ValidateOperationRequestResource) (result OperationValidateOperationResponse, err error) {
	req, err := c.preparerForOperationValidate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "OperationValidate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "OperationValidate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForOperationValidate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "OperationValidate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForOperationValidate prepares the OperationValidate request.
func (c BmsClient) preparerForOperationValidate(ctx context.Context, id VaultId, input ValidateOperationRequestResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backupValidateOperation", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForOperationValidate handles the response to the OperationValidate request. The method always
// closes the http.Response Body.
func (c BmsClient) responderForOperationValidate(resp *http.Response) (result OperationValidateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

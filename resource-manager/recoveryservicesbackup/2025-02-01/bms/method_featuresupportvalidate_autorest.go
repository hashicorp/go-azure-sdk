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

type FeatureSupportValidateOperationResponse struct {
	HttpResponse *http.Response
	Model        *AzureVMResourceFeatureSupportResponse
}

// FeatureSupportValidate ...
func (c BmsClient) FeatureSupportValidate(ctx context.Context, id LocationId, input FeatureSupportRequest) (result FeatureSupportValidateOperationResponse, err error) {
	req, err := c.preparerForFeatureSupportValidate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "FeatureSupportValidate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "FeatureSupportValidate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForFeatureSupportValidate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bms.BmsClient", "FeatureSupportValidate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForFeatureSupportValidate prepares the FeatureSupportValidate request.
func (c BmsClient) preparerForFeatureSupportValidate(ctx context.Context, id LocationId, input FeatureSupportRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backupValidateFeatures", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForFeatureSupportValidate handles the response to the FeatureSupportValidate request. The method always
// closes the http.Response Body.
func (c BmsClient) responderForFeatureSupportValidate(resp *http.Response) (result FeatureSupportValidateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

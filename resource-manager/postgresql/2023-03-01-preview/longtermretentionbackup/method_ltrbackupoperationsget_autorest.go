package longtermretentionbackup

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LtrBackupOperationsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *LtrServerBackupOperation
}

// LtrBackupOperationsGet ...
func (c LongTermRetentionBackupClient) LtrBackupOperationsGet(ctx context.Context, id LtrBackupOperationId) (result LtrBackupOperationsGetOperationResponse, err error) {
	req, err := c.preparerForLtrBackupOperationsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "longtermretentionbackup.LongTermRetentionBackupClient", "LtrBackupOperationsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "longtermretentionbackup.LongTermRetentionBackupClient", "LtrBackupOperationsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForLtrBackupOperationsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "longtermretentionbackup.LongTermRetentionBackupClient", "LtrBackupOperationsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForLtrBackupOperationsGet prepares the LtrBackupOperationsGet request.
func (c LongTermRetentionBackupClient) preparerForLtrBackupOperationsGet(ctx context.Context, id LtrBackupOperationId) (*http.Request, error) {
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

// responderForLtrBackupOperationsGet handles the response to the LtrBackupOperationsGet request. The method always
// closes the http.Response Body.
func (c LongTermRetentionBackupClient) responderForLtrBackupOperationsGet(resp *http.Response) (result LtrBackupOperationsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

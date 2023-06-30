package feature

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *BatchFeatureStatus
}

// AccountGet ...
func (c FeatureClient) AccountGet(ctx context.Context, id AccountId, input BatchFeatureRequest) (result AccountGetOperationResponse, err error) {
	req, err := c.preparerForAccountGet(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "feature.FeatureClient", "AccountGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "feature.FeatureClient", "AccountGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAccountGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "feature.FeatureClient", "AccountGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAccountGet prepares the AccountGet request.
func (c FeatureClient) preparerForAccountGet(ctx context.Context, id AccountId, input BatchFeatureRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listFeatures", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAccountGet handles the response to the AccountGet request. The method always
// closes the http.Response Body.
func (c FeatureClient) responderForAccountGet(resp *http.Response) (result AccountGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

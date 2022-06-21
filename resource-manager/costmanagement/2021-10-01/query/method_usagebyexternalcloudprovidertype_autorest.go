package query

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageByExternalCloudProviderTypeOperationResponse struct {
	HttpResponse *http.Response
	Model        *QueryResult
}

// UsageByExternalCloudProviderType ...
func (c QueryClient) UsageByExternalCloudProviderType(ctx context.Context, id ExternalCloudProviderTypeId, input QueryDefinition) (result UsageByExternalCloudProviderTypeOperationResponse, err error) {
	req, err := c.preparerForUsageByExternalCloudProviderType(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "query.QueryClient", "UsageByExternalCloudProviderType", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "query.QueryClient", "UsageByExternalCloudProviderType", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUsageByExternalCloudProviderType(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "query.QueryClient", "UsageByExternalCloudProviderType", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUsageByExternalCloudProviderType prepares the UsageByExternalCloudProviderType request.
func (c QueryClient) preparerForUsageByExternalCloudProviderType(ctx context.Context, id ExternalCloudProviderTypeId, input QueryDefinition) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/query", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUsageByExternalCloudProviderType handles the response to the UsageByExternalCloudProviderType request. The method always
// closes the http.Response Body.
func (c QueryClient) responderForUsageByExternalCloudProviderType(resp *http.Response) (result UsageByExternalCloudProviderTypeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}

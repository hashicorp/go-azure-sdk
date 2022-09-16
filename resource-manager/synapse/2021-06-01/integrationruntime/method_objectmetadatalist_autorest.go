package integrationruntime

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectMetadataListOperationResponse struct {
	HttpResponse *http.Response
	Model        *SsisObjectMetadataListResponse
}

// ObjectMetadataList ...
func (c IntegrationRuntimeClient) ObjectMetadataList(ctx context.Context, id IntegrationRuntimeId, input GetSsisObjectMetadataRequest) (result ObjectMetadataListOperationResponse, err error) {
	req, err := c.preparerForObjectMetadataList(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "ObjectMetadataList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "ObjectMetadataList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForObjectMetadataList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "ObjectMetadataList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForObjectMetadataList prepares the ObjectMetadataList request.
func (c IntegrationRuntimeClient) preparerForObjectMetadataList(ctx context.Context, id IntegrationRuntimeId, input GetSsisObjectMetadataRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getObjectMetadata", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForObjectMetadataList handles the response to the ObjectMetadataList request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimeClient) responderForObjectMetadataList(resp *http.Response) (result ObjectMetadataListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

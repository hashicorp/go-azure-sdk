package dimensions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ByExternalCloudProviderTypeOperationResponse struct {
	HttpResponse *http.Response
	Model        *Dimension
}

type ByExternalCloudProviderTypeOperationOptions struct {
	Expand *string
	Filter *string
	Top    *int64
}

func DefaultByExternalCloudProviderTypeOperationOptions() ByExternalCloudProviderTypeOperationOptions {
	return ByExternalCloudProviderTypeOperationOptions{}
}

func (o ByExternalCloudProviderTypeOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ByExternalCloudProviderTypeOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ByExternalCloudProviderType ...
func (c DimensionsClient) ByExternalCloudProviderType(ctx context.Context, id ExternalCloudProviderTypeId, options ByExternalCloudProviderTypeOperationOptions) (result ByExternalCloudProviderTypeOperationResponse, err error) {
	req, err := c.preparerForByExternalCloudProviderType(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dimensions.DimensionsClient", "ByExternalCloudProviderType", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "dimensions.DimensionsClient", "ByExternalCloudProviderType", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForByExternalCloudProviderType(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dimensions.DimensionsClient", "ByExternalCloudProviderType", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForByExternalCloudProviderType prepares the ByExternalCloudProviderType request.
func (c DimensionsClient) preparerForByExternalCloudProviderType(ctx context.Context, id ExternalCloudProviderTypeId, options ByExternalCloudProviderTypeOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/dimensions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForByExternalCloudProviderType handles the response to the ByExternalCloudProviderType request. The method always
// closes the http.Response Body.
func (c DimensionsClient) responderForByExternalCloudProviderType(resp *http.Response) (result ByExternalCloudProviderTypeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

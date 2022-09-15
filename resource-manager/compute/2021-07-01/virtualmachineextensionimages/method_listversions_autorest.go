package virtualmachineextensionimages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListVersionsOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VirtualMachineExtensionImage
}

type ListVersionsOperationOptions struct {
	Filter  *string
	Orderby *string
	Top     *int64
}

func DefaultListVersionsOperationOptions() ListVersionsOperationOptions {
	return ListVersionsOperationOptions{}
}

func (o ListVersionsOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o ListVersionsOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	if o.Orderby != nil {
		out["$orderby"] = *o.Orderby
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// ListVersions ...
func (c VirtualMachineExtensionImagesClient) ListVersions(ctx context.Context, id TypeId, options ListVersionsOperationOptions) (result ListVersionsOperationResponse, err error) {
	req, err := c.preparerForListVersions(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineextensionimages.VirtualMachineExtensionImagesClient", "ListVersions", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineextensionimages.VirtualMachineExtensionImagesClient", "ListVersions", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListVersions(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineextensionimages.VirtualMachineExtensionImagesClient", "ListVersions", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListVersions prepares the ListVersions request.
func (c VirtualMachineExtensionImagesClient) preparerForListVersions(ctx context.Context, id TypeId, options ListVersionsOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/versions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListVersions handles the response to the ListVersions request. The method always
// closes the http.Response Body.
func (c VirtualMachineExtensionImagesClient) responderForListVersions(resp *http.Response) (result ListVersionsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

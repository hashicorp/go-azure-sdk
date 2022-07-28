package virtualmachineimages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgeZoneListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VirtualMachineImageResource
}

type EdgeZoneListOperationOptions struct {
	Expand  *string
	Orderby *string
	Top     *int64
}

func DefaultEdgeZoneListOperationOptions() EdgeZoneListOperationOptions {
	return EdgeZoneListOperationOptions{}
}

func (o EdgeZoneListOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o EdgeZoneListOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Expand != nil {
		out["$expand"] = *o.Expand
	}

	if o.Orderby != nil {
		out["$orderby"] = *o.Orderby
	}

	if o.Top != nil {
		out["$top"] = *o.Top
	}

	return out
}

// EdgeZoneList ...
func (c VirtualMachineImagesClient) EdgeZoneList(ctx context.Context, id OfferSkuId, options EdgeZoneListOperationOptions) (result EdgeZoneListOperationResponse, err error) {
	req, err := c.preparerForEdgeZoneList(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneList", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneList", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForEdgeZoneList(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachineimages.VirtualMachineImagesClient", "EdgeZoneList", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForEdgeZoneList prepares the EdgeZoneList request.
func (c VirtualMachineImagesClient) preparerForEdgeZoneList(ctx context.Context, id OfferSkuId, options EdgeZoneListOperationOptions) (*http.Request, error) {
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

// responderForEdgeZoneList handles the response to the EdgeZoneList request. The method always
// closes the http.Response Body.
func (c VirtualMachineImagesClient) responderForEdgeZoneList(resp *http.Response) (result EdgeZoneListOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

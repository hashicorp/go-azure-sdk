package aggregatedcost

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetByManagementGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *ManagementGroupAggregatedCostResult
}

type GetByManagementGroupOperationOptions struct {
	Filter *string
}

func DefaultGetByManagementGroupOperationOptions() GetByManagementGroupOperationOptions {
	return GetByManagementGroupOperationOptions{}
}

func (o GetByManagementGroupOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

func (o GetByManagementGroupOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	if o.Filter != nil {
		out["$filter"] = *o.Filter
	}

	return out
}

// GetByManagementGroup ...
func (c AggregatedCostClient) GetByManagementGroup(ctx context.Context, id commonids.ManagementGroupId, options GetByManagementGroupOperationOptions) (result GetByManagementGroupOperationResponse, err error) {
	req, err := c.preparerForGetByManagementGroup(ctx, id, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aggregatedcost.AggregatedCostClient", "GetByManagementGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "aggregatedcost.AggregatedCostClient", "GetByManagementGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetByManagementGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "aggregatedcost.AggregatedCostClient", "GetByManagementGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetByManagementGroup prepares the GetByManagementGroup request.
func (c AggregatedCostClient) preparerForGetByManagementGroup(ctx context.Context, id commonids.ManagementGroupId, options GetByManagementGroupOperationOptions) (*http.Request, error) {
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
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.Consumption/aggregatedcost", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetByManagementGroup handles the response to the GetByManagementGroup request. The method always
// closes the http.Response Body.
func (c AggregatedCostClient) responderForGetByManagementGroup(resp *http.Response) (result GetByManagementGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}

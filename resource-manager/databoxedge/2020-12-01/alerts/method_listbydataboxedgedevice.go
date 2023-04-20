package alerts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDataBoxEdgeDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Alert
}

type ListByDataBoxEdgeDeviceCompleteResult struct {
	Items []Alert
}

// ListByDataBoxEdgeDevice ...
func (c AlertsClient) ListByDataBoxEdgeDevice(ctx context.Context, id DataBoxEdgeDeviceId) (result ListByDataBoxEdgeDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/alerts", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]Alert `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByDataBoxEdgeDeviceComplete retrieves all the results into a single object
func (c AlertsClient) ListByDataBoxEdgeDeviceComplete(ctx context.Context, id DataBoxEdgeDeviceId) (ListByDataBoxEdgeDeviceCompleteResult, error) {
	return c.ListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx, id, AlertOperationPredicate{})
}

// ListByDataBoxEdgeDeviceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AlertsClient) ListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx context.Context, id DataBoxEdgeDeviceId, predicate AlertOperationPredicate) (result ListByDataBoxEdgeDeviceCompleteResult, err error) {
	items := make([]Alert, 0)

	resp, err := c.ListByDataBoxEdgeDevice(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListByDataBoxEdgeDeviceCompleteResult{
		Items: items,
	}
	return
}

package managedinstances

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByInstancePoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ManagedInstance
}

type ListByInstancePoolCompleteResult struct {
	Items []ManagedInstance
}

// ListByInstancePool ...
func (c ManagedInstancesClient) ListByInstancePool(ctx context.Context, id InstancePoolId) (result ListByInstancePoolOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/managedInstances", id.ID()),
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
		Values *[]ManagedInstance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByInstancePoolComplete retrieves all the results into a single object
func (c ManagedInstancesClient) ListByInstancePoolComplete(ctx context.Context, id InstancePoolId) (ListByInstancePoolCompleteResult, error) {
	return c.ListByInstancePoolCompleteMatchingPredicate(ctx, id, ManagedInstanceOperationPredicate{})
}

// ListByInstancePoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ManagedInstancesClient) ListByInstancePoolCompleteMatchingPredicate(ctx context.Context, id InstancePoolId, predicate ManagedInstanceOperationPredicate) (result ListByInstancePoolCompleteResult, err error) {
	items := make([]ManagedInstance, 0)

	resp, err := c.ListByInstancePool(ctx, id)
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

	result = ListByInstancePoolCompleteResult{
		Items: items,
	}
	return
}

package sqlpoolsworkloadgroups

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolWorkloadGroupListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadGroup
}

type SqlPoolWorkloadGroupListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadGroup
}

// SqlPoolWorkloadGroupList ...
func (c SqlPoolsWorkloadGroupsClient) SqlPoolWorkloadGroupList(ctx context.Context, id SqlPoolId) (result SqlPoolWorkloadGroupListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/workloadGroups", id.ID()),
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
		Values *[]WorkloadGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SqlPoolWorkloadGroupListComplete retrieves all the results into a single object
func (c SqlPoolsWorkloadGroupsClient) SqlPoolWorkloadGroupListComplete(ctx context.Context, id SqlPoolId) (SqlPoolWorkloadGroupListCompleteResult, error) {
	return c.SqlPoolWorkloadGroupListCompleteMatchingPredicate(ctx, id, WorkloadGroupOperationPredicate{})
}

// SqlPoolWorkloadGroupListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SqlPoolsWorkloadGroupsClient) SqlPoolWorkloadGroupListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate WorkloadGroupOperationPredicate) (result SqlPoolWorkloadGroupListCompleteResult, err error) {
	items := make([]WorkloadGroup, 0)

	resp, err := c.SqlPoolWorkloadGroupList(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = SqlPoolWorkloadGroupListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

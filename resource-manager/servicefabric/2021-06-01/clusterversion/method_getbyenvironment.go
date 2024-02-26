package clusterversion

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetByEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ClusterCodeVersionsResult
}

type GetByEnvironmentCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ClusterCodeVersionsResult
}

// GetByEnvironment ...
func (c ClusterVersionClient) GetByEnvironment(ctx context.Context, id EnvironmentClusterVersionId) (result GetByEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       id.ID(),
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
		Values *[]ClusterCodeVersionsResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetByEnvironmentComplete retrieves all the results into a single object
func (c ClusterVersionClient) GetByEnvironmentComplete(ctx context.Context, id EnvironmentClusterVersionId) (GetByEnvironmentCompleteResult, error) {
	return c.GetByEnvironmentCompleteMatchingPredicate(ctx, id, ClusterCodeVersionsResultOperationPredicate{})
}

// GetByEnvironmentCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ClusterVersionClient) GetByEnvironmentCompleteMatchingPredicate(ctx context.Context, id EnvironmentClusterVersionId, predicate ClusterCodeVersionsResultOperationPredicate) (result GetByEnvironmentCompleteResult, err error) {
	items := make([]ClusterCodeVersionsResult, 0)

	resp, err := c.GetByEnvironment(ctx, id)
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

	result = GetByEnvironmentCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

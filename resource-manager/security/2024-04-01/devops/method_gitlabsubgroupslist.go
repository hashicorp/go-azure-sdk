package devops

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GitLabSubgroupsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GitLabGroup
}

type GitLabSubgroupsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GitLabGroup
}

// GitLabSubgroupsList ...
func (c DevOpsClient) GitLabSubgroupsList(ctx context.Context, id GitLabGroupId) (result GitLabSubgroupsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/listSubgroups", id.ID()),
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
		Values *[]GitLabGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GitLabSubgroupsListComplete retrieves all the results into a single object
func (c DevOpsClient) GitLabSubgroupsListComplete(ctx context.Context, id GitLabGroupId) (GitLabSubgroupsListCompleteResult, error) {
	return c.GitLabSubgroupsListCompleteMatchingPredicate(ctx, id, GitLabGroupOperationPredicate{})
}

// GitLabSubgroupsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DevOpsClient) GitLabSubgroupsListCompleteMatchingPredicate(ctx context.Context, id GitLabGroupId, predicate GitLabGroupOperationPredicate) (result GitLabSubgroupsListCompleteResult, err error) {
	items := make([]GitLabGroup, 0)

	resp, err := c.GitLabSubgroupsList(ctx, id)
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

	result = GitLabSubgroupsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

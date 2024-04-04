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

type GitLabProjectsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GitLabProject
}

type GitLabProjectsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GitLabProject
}

// GitLabProjectsList ...
func (c DevOpsClient) GitLabProjectsList(ctx context.Context, id GitLabGroupId) (result GitLabProjectsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/projects", id.ID()),
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
		Values *[]GitLabProject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GitLabProjectsListComplete retrieves all the results into a single object
func (c DevOpsClient) GitLabProjectsListComplete(ctx context.Context, id GitLabGroupId) (GitLabProjectsListCompleteResult, error) {
	return c.GitLabProjectsListCompleteMatchingPredicate(ctx, id, GitLabProjectOperationPredicate{})
}

// GitLabProjectsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DevOpsClient) GitLabProjectsListCompleteMatchingPredicate(ctx context.Context, id GitLabGroupId, predicate GitLabProjectOperationPredicate) (result GitLabProjectsListCompleteResult, err error) {
	items := make([]GitLabProject, 0)

	resp, err := c.GitLabProjectsList(ctx, id)
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

	result = GitLabProjectsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

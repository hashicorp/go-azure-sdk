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

type GitLabGroupsListAvailableOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GitLabGroup
}

type GitLabGroupsListAvailableCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GitLabGroup
}

// GitLabGroupsListAvailable ...
func (c DevOpsClient) GitLabGroupsListAvailable(ctx context.Context, id SecurityConnectorId) (result GitLabGroupsListAvailableOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/devops/default/listAvailableGitLabGroups", id.ID()),
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

// GitLabGroupsListAvailableComplete retrieves all the results into a single object
func (c DevOpsClient) GitLabGroupsListAvailableComplete(ctx context.Context, id SecurityConnectorId) (GitLabGroupsListAvailableCompleteResult, error) {
	return c.GitLabGroupsListAvailableCompleteMatchingPredicate(ctx, id, GitLabGroupOperationPredicate{})
}

// GitLabGroupsListAvailableCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DevOpsClient) GitLabGroupsListAvailableCompleteMatchingPredicate(ctx context.Context, id SecurityConnectorId, predicate GitLabGroupOperationPredicate) (result GitLabGroupsListAvailableCompleteResult, err error) {
	items := make([]GitLabGroup, 0)

	resp, err := c.GitLabGroupsListAvailable(ctx, id)
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

	result = GitLabGroupsListAvailableCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

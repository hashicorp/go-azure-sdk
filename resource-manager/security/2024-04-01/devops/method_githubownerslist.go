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

type GitHubOwnersListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GitHubOwner
}

type GitHubOwnersListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GitHubOwner
}

// GitHubOwnersList ...
func (c DevOpsClient) GitHubOwnersList(ctx context.Context, id SecurityConnectorId) (result GitHubOwnersListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/devops/default/gitHubOwners", id.ID()),
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
		Values *[]GitHubOwner `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GitHubOwnersListComplete retrieves all the results into a single object
func (c DevOpsClient) GitHubOwnersListComplete(ctx context.Context, id SecurityConnectorId) (GitHubOwnersListCompleteResult, error) {
	return c.GitHubOwnersListCompleteMatchingPredicate(ctx, id, GitHubOwnerOperationPredicate{})
}

// GitHubOwnersListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DevOpsClient) GitHubOwnersListCompleteMatchingPredicate(ctx context.Context, id SecurityConnectorId, predicate GitHubOwnerOperationPredicate) (result GitHubOwnersListCompleteResult, err error) {
	items := make([]GitHubOwner, 0)

	resp, err := c.GitHubOwnersList(ctx, id)
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

	result = GitHubOwnersListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

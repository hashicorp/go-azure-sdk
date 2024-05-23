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

type GitHubOwnersListAvailableOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GitHubOwner
}

type GitHubOwnersListAvailableCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GitHubOwner
}

// GitHubOwnersListAvailable ...
func (c DevOpsClient) GitHubOwnersListAvailable(ctx context.Context, id SecurityConnectorId) (result GitHubOwnersListAvailableOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/devops/default/listAvailableGitHubOwners", id.ID()),
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

// GitHubOwnersListAvailableComplete retrieves all the results into a single object
func (c DevOpsClient) GitHubOwnersListAvailableComplete(ctx context.Context, id SecurityConnectorId) (GitHubOwnersListAvailableCompleteResult, error) {
	return c.GitHubOwnersListAvailableCompleteMatchingPredicate(ctx, id, GitHubOwnerOperationPredicate{})
}

// GitHubOwnersListAvailableCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DevOpsClient) GitHubOwnersListAvailableCompleteMatchingPredicate(ctx context.Context, id SecurityConnectorId, predicate GitHubOwnerOperationPredicate) (result GitHubOwnersListAvailableCompleteResult, err error) {
	items := make([]GitHubOwner, 0)

	resp, err := c.GitHubOwnersListAvailable(ctx, id)
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

	result = GitHubOwnersListAvailableCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

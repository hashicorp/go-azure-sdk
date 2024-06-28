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

type GitHubReposListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]GitHubRepository
}

type GitHubReposListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []GitHubRepository
}

type GitHubReposListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GitHubReposListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GitHubReposList ...
func (c DevOpsClient) GitHubReposList(ctx context.Context, id GitHubOwnerId) (result GitHubReposListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &GitHubReposListCustomPager{},
		Path:       fmt.Sprintf("%s/repos", id.ID()),
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
		Values *[]GitHubRepository `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GitHubReposListComplete retrieves all the results into a single object
func (c DevOpsClient) GitHubReposListComplete(ctx context.Context, id GitHubOwnerId) (GitHubReposListCompleteResult, error) {
	return c.GitHubReposListCompleteMatchingPredicate(ctx, id, GitHubRepositoryOperationPredicate{})
}

// GitHubReposListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DevOpsClient) GitHubReposListCompleteMatchingPredicate(ctx context.Context, id GitHubOwnerId, predicate GitHubRepositoryOperationPredicate) (result GitHubReposListCompleteResult, err error) {
	items := make([]GitHubRepository, 0)

	resp, err := c.GitHubReposList(ctx, id)
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

	result = GitHubReposListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

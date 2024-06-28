package repositories

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControllistRepositoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Repo
}

type SourceControllistRepositoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Repo
}

type SourceControllistRepositoriesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *SourceControllistRepositoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// SourceControllistRepositories ...
func (c RepositoriesClient) SourceControllistRepositories(ctx context.Context, id WorkspaceId, input RepoType) (result SourceControllistRepositoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &SourceControllistRepositoriesCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/listRepositories", id.ID()),
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
		Values *[]Repo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// SourceControllistRepositoriesComplete retrieves all the results into a single object
func (c RepositoriesClient) SourceControllistRepositoriesComplete(ctx context.Context, id WorkspaceId, input RepoType) (SourceControllistRepositoriesCompleteResult, error) {
	return c.SourceControllistRepositoriesCompleteMatchingPredicate(ctx, id, input, RepoOperationPredicate{})
}

// SourceControllistRepositoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RepositoriesClient) SourceControllistRepositoriesCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, input RepoType, predicate RepoOperationPredicate) (result SourceControllistRepositoriesCompleteResult, err error) {
	items := make([]Repo, 0)

	resp, err := c.SourceControllistRepositories(ctx, id, input)
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

	result = SourceControllistRepositoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

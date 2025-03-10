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

type AzureDevOpsReposListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AzureDevOpsRepository
}

type AzureDevOpsReposListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AzureDevOpsRepository
}

type AzureDevOpsReposListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AzureDevOpsReposListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AzureDevOpsReposList ...
func (c DevOpsClient) AzureDevOpsReposList(ctx context.Context, id ProjectId) (result AzureDevOpsReposListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &AzureDevOpsReposListCustomPager{},
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
		Values *[]AzureDevOpsRepository `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AzureDevOpsReposListComplete retrieves all the results into a single object
func (c DevOpsClient) AzureDevOpsReposListComplete(ctx context.Context, id ProjectId) (AzureDevOpsReposListCompleteResult, error) {
	return c.AzureDevOpsReposListCompleteMatchingPredicate(ctx, id, AzureDevOpsRepositoryOperationPredicate{})
}

// AzureDevOpsReposListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DevOpsClient) AzureDevOpsReposListCompleteMatchingPredicate(ctx context.Context, id ProjectId, predicate AzureDevOpsRepositoryOperationPredicate) (result AzureDevOpsReposListCompleteResult, err error) {
	items := make([]AzureDevOpsRepository, 0)

	resp, err := c.AzureDevOpsReposList(ctx, id)
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

	result = AzureDevOpsReposListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

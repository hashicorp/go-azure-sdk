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

type ListByEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ClusterCodeVersionsResult
}

type ListByEnvironmentCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ClusterCodeVersionsResult
}

type ListByEnvironmentCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByEnvironmentCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByEnvironment ...
func (c ClusterVersionClient) ListByEnvironment(ctx context.Context, id EnvironmentId) (result ListByEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByEnvironmentCustomPager{},
		Path:       fmt.Sprintf("%s/clusterVersions", id.ID()),
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

// ListByEnvironmentComplete retrieves all the results into a single object
func (c ClusterVersionClient) ListByEnvironmentComplete(ctx context.Context, id EnvironmentId) (ListByEnvironmentCompleteResult, error) {
	return c.ListByEnvironmentCompleteMatchingPredicate(ctx, id, ClusterCodeVersionsResultOperationPredicate{})
}

// ListByEnvironmentCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ClusterVersionClient) ListByEnvironmentCompleteMatchingPredicate(ctx context.Context, id EnvironmentId, predicate ClusterCodeVersionsResultOperationPredicate) (result ListByEnvironmentCompleteResult, err error) {
	items := make([]ClusterCodeVersionsResult, 0)

	resp, err := c.ListByEnvironment(ctx, id)
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

	result = ListByEnvironmentCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

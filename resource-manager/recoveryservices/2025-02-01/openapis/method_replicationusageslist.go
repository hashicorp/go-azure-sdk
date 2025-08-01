package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationUsagesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ReplicationUsage
}

type ReplicationUsagesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ReplicationUsage
}

type ReplicationUsagesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ReplicationUsagesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ReplicationUsagesList ...
func (c OpenapisClient) ReplicationUsagesList(ctx context.Context, id VaultId) (result ReplicationUsagesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ReplicationUsagesListCustomPager{},
		Path:       fmt.Sprintf("%s/replicationUsages", id.ID()),
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
		Values *[]ReplicationUsage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ReplicationUsagesListComplete retrieves all the results into a single object
func (c OpenapisClient) ReplicationUsagesListComplete(ctx context.Context, id VaultId) (ReplicationUsagesListCompleteResult, error) {
	return c.ReplicationUsagesListCompleteMatchingPredicate(ctx, id, ReplicationUsageOperationPredicate{})
}

// ReplicationUsagesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) ReplicationUsagesListCompleteMatchingPredicate(ctx context.Context, id VaultId, predicate ReplicationUsageOperationPredicate) (result ReplicationUsagesListCompleteResult, err error) {
	items := make([]ReplicationUsage, 0)

	resp, err := c.ReplicationUsagesList(ctx, id)
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

	result = ReplicationUsagesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

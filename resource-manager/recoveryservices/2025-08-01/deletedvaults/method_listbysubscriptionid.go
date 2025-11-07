package deletedvaults

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBySubscriptionIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeletedVault
}

type ListBySubscriptionIdCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeletedVault
}

type ListBySubscriptionIdCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListBySubscriptionIdCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListBySubscriptionId ...
func (c DeletedVaultsClient) ListBySubscriptionId(ctx context.Context, id LocationId) (result ListBySubscriptionIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListBySubscriptionIdCustomPager{},
		Path:       fmt.Sprintf("%s/deletedVaults", id.ID()),
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
		Values *[]DeletedVault `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListBySubscriptionIdComplete retrieves all the results into a single object
func (c DeletedVaultsClient) ListBySubscriptionIdComplete(ctx context.Context, id LocationId) (ListBySubscriptionIdCompleteResult, error) {
	return c.ListBySubscriptionIdCompleteMatchingPredicate(ctx, id, DeletedVaultOperationPredicate{})
}

// ListBySubscriptionIdCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeletedVaultsClient) ListBySubscriptionIdCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate DeletedVaultOperationPredicate) (result ListBySubscriptionIdCompleteResult, err error) {
	items := make([]DeletedVault, 0)

	resp, err := c.ListBySubscriptionId(ctx, id)
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

	result = ListBySubscriptionIdCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

package scopedrolememberof

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListScopedRoleMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ScopedRoleMembership
}

type ListScopedRoleMemberOfCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ScopedRoleMembership
}

type ListScopedRoleMemberOfCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListScopedRoleMemberOfCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListScopedRoleMemberOf ...
func (c ScopedRoleMemberOfClient) ListScopedRoleMemberOf(ctx context.Context) (result ListScopedRoleMemberOfOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListScopedRoleMemberOfCustomPager{},
		Path:       "/me/scopedRoleMemberOf",
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
		Values *[]beta.ScopedRoleMembership `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListScopedRoleMemberOfComplete retrieves all the results into a single object
func (c ScopedRoleMemberOfClient) ListScopedRoleMemberOfComplete(ctx context.Context) (ListScopedRoleMemberOfCompleteResult, error) {
	return c.ListScopedRoleMemberOfCompleteMatchingPredicate(ctx, ScopedRoleMembershipOperationPredicate{})
}

// ListScopedRoleMemberOfCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ScopedRoleMemberOfClient) ListScopedRoleMemberOfCompleteMatchingPredicate(ctx context.Context, predicate ScopedRoleMembershipOperationPredicate) (result ListScopedRoleMemberOfCompleteResult, err error) {
	items := make([]beta.ScopedRoleMembership, 0)

	resp, err := c.ListScopedRoleMemberOf(ctx)
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

	result = ListScopedRoleMemberOfCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

package profileaward

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

type ListProfileAwardsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PersonAward
}

type ListProfileAwardsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PersonAward
}

type ListProfileAwardsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileAwardsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileAwards ...
func (c ProfileAwardClient) ListProfileAwards(ctx context.Context) (result ListProfileAwardsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfileAwardsCustomPager{},
		Path:       "/me/profile/awards",
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
		Values *[]beta.PersonAward `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileAwardsComplete retrieves all the results into a single object
func (c ProfileAwardClient) ListProfileAwardsComplete(ctx context.Context) (ListProfileAwardsCompleteResult, error) {
	return c.ListProfileAwardsCompleteMatchingPredicate(ctx, PersonAwardOperationPredicate{})
}

// ListProfileAwardsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileAwardClient) ListProfileAwardsCompleteMatchingPredicate(ctx context.Context, predicate PersonAwardOperationPredicate) (result ListProfileAwardsCompleteResult, err error) {
	items := make([]beta.PersonAward, 0)

	resp, err := c.ListProfileAwards(ctx)
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

	result = ListProfileAwardsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

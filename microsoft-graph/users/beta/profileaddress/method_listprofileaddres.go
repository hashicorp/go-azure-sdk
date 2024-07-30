package profileaddress

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

type ListProfileAddresOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ItemAddress
}

type ListProfileAddresCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ItemAddress
}

type ListProfileAddresCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileAddresCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileAddres ...
func (c ProfileAddressClient) ListProfileAddres(ctx context.Context, id UserId) (result ListProfileAddresOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfileAddresCustomPager{},
		Path:       fmt.Sprintf("%s/profile/addresses", id.ID()),
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
		Values *[]beta.ItemAddress `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileAddresComplete retrieves all the results into a single object
func (c ProfileAddressClient) ListProfileAddresComplete(ctx context.Context, id UserId) (ListProfileAddresCompleteResult, error) {
	return c.ListProfileAddresCompleteMatchingPredicate(ctx, id, ItemAddressOperationPredicate{})
}

// ListProfileAddresCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileAddressClient) ListProfileAddresCompleteMatchingPredicate(ctx context.Context, id UserId, predicate ItemAddressOperationPredicate) (result ListProfileAddresCompleteResult, err error) {
	items := make([]beta.ItemAddress, 0)

	resp, err := c.ListProfileAddres(ctx, id)
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

	result = ListProfileAddresCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

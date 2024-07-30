package virtualeventwebinar

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

type ListVirtualEventWebinarsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.VirtualEventWebinar
}

type ListVirtualEventWebinarsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.VirtualEventWebinar
}

type ListVirtualEventWebinarsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEventWebinarsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEventWebinars ...
func (c VirtualEventWebinarClient) ListVirtualEventWebinars(ctx context.Context) (result ListVirtualEventWebinarsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListVirtualEventWebinarsCustomPager{},
		Path:       "/me/virtualEvents/webinars",
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
		Values *[]beta.VirtualEventWebinar `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEventWebinarsComplete retrieves all the results into a single object
func (c VirtualEventWebinarClient) ListVirtualEventWebinarsComplete(ctx context.Context) (ListVirtualEventWebinarsCompleteResult, error) {
	return c.ListVirtualEventWebinarsCompleteMatchingPredicate(ctx, VirtualEventWebinarOperationPredicate{})
}

// ListVirtualEventWebinarsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEventWebinarClient) ListVirtualEventWebinarsCompleteMatchingPredicate(ctx context.Context, predicate VirtualEventWebinarOperationPredicate) (result ListVirtualEventWebinarsCompleteResult, err error) {
	items := make([]beta.VirtualEventWebinar, 0)

	resp, err := c.ListVirtualEventWebinars(ctx)
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

	result = ListVirtualEventWebinarsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

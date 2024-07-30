package profilepatent

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

type ListProfilePatentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ItemPatent
}

type ListProfilePatentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ItemPatent
}

type ListProfilePatentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfilePatentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfilePatents ...
func (c ProfilePatentClient) ListProfilePatents(ctx context.Context) (result ListProfilePatentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfilePatentsCustomPager{},
		Path:       "/me/profile/patents",
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
		Values *[]beta.ItemPatent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfilePatentsComplete retrieves all the results into a single object
func (c ProfilePatentClient) ListProfilePatentsComplete(ctx context.Context) (ListProfilePatentsCompleteResult, error) {
	return c.ListProfilePatentsCompleteMatchingPredicate(ctx, ItemPatentOperationPredicate{})
}

// ListProfilePatentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfilePatentClient) ListProfilePatentsCompleteMatchingPredicate(ctx context.Context, predicate ItemPatentOperationPredicate) (result ListProfilePatentsCompleteResult, err error) {
	items := make([]beta.ItemPatent, 0)

	resp, err := c.ListProfilePatents(ctx)
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

	result = ListProfilePatentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

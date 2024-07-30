package transitivememberof

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListTransitiveMemberOfOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListTransitiveMemberOfCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListTransitiveMemberOfCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTransitiveMemberOfCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTransitiveMemberOf ...
func (c TransitiveMemberOfClient) ListTransitiveMemberOf(ctx context.Context, id UserId) (result ListTransitiveMemberOfOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTransitiveMemberOfCustomPager{},
		Path:       fmt.Sprintf("%s/transitiveMemberOf", id.ID()),
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
		Values *[]stable.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTransitiveMemberOfComplete retrieves all the results into a single object
func (c TransitiveMemberOfClient) ListTransitiveMemberOfComplete(ctx context.Context, id UserId) (ListTransitiveMemberOfCompleteResult, error) {
	return c.ListTransitiveMemberOfCompleteMatchingPredicate(ctx, id, DirectoryObjectOperationPredicate{})
}

// ListTransitiveMemberOfCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TransitiveMemberOfClient) ListTransitiveMemberOfCompleteMatchingPredicate(ctx context.Context, id UserId, predicate DirectoryObjectOperationPredicate) (result ListTransitiveMemberOfCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListTransitiveMemberOf(ctx, id)
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

	result = ListTransitiveMemberOfCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

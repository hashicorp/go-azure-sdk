package transitivemember

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

type ListTransitiveMembersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DirectoryObject
}

type ListTransitiveMembersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DirectoryObject
}

type ListTransitiveMembersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTransitiveMembersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTransitiveMembers ...
func (c TransitiveMemberClient) ListTransitiveMembers(ctx context.Context, id GroupId) (result ListTransitiveMembersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTransitiveMembersCustomPager{},
		Path:       fmt.Sprintf("%s/transitiveMembers", id.ID()),
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

// ListTransitiveMembersComplete retrieves all the results into a single object
func (c TransitiveMemberClient) ListTransitiveMembersComplete(ctx context.Context, id GroupId) (ListTransitiveMembersCompleteResult, error) {
	return c.ListTransitiveMembersCompleteMatchingPredicate(ctx, id, DirectoryObjectOperationPredicate{})
}

// ListTransitiveMembersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TransitiveMemberClient) ListTransitiveMembersCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate DirectoryObjectOperationPredicate) (result ListTransitiveMembersCompleteResult, err error) {
	items := make([]stable.DirectoryObject, 0)

	resp, err := c.ListTransitiveMembers(ctx, id)
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

	result = ListTransitiveMembersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

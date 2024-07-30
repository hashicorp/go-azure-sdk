package deleteditem

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

type GetDirectoryDeletedItemsByIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type GetDirectoryDeletedItemsByIdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type GetDirectoryDeletedItemsByIdsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetDirectoryDeletedItemsByIdsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetDirectoryDeletedItemsByIds ...
func (c DeletedItemClient) GetDirectoryDeletedItemsByIds(ctx context.Context, input GetDirectoryDeletedItemsByIdsRequest) (result GetDirectoryDeletedItemsByIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &GetDirectoryDeletedItemsByIdsCustomPager{},
		Path:       "/directory/deletedItems/getByIds",
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
		Values *[]beta.DirectoryObject `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetDirectoryDeletedItemsByIdsComplete retrieves all the results into a single object
func (c DeletedItemClient) GetDirectoryDeletedItemsByIdsComplete(ctx context.Context, input GetDirectoryDeletedItemsByIdsRequest) (GetDirectoryDeletedItemsByIdsCompleteResult, error) {
	return c.GetDirectoryDeletedItemsByIdsCompleteMatchingPredicate(ctx, input, DirectoryObjectOperationPredicate{})
}

// GetDirectoryDeletedItemsByIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeletedItemClient) GetDirectoryDeletedItemsByIdsCompleteMatchingPredicate(ctx context.Context, input GetDirectoryDeletedItemsByIdsRequest, predicate DirectoryObjectOperationPredicate) (result GetDirectoryDeletedItemsByIdsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.GetDirectoryDeletedItemsByIds(ctx, input)
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

	result = GetDirectoryDeletedItemsByIdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

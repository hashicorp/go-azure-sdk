package user

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

type GetUsersByIdsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryObject
}

type GetUsersByIdsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryObject
}

type GetUsersByIdsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetUsersByIdsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetUsersByIds ...
func (c UserClient) GetUsersByIds(ctx context.Context, input GetUsersByIdsRequest) (result GetUsersByIdsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &GetUsersByIdsCustomPager{},
		Path:       "/users/getByIds",
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

// GetUsersByIdsComplete retrieves all the results into a single object
func (c UserClient) GetUsersByIdsComplete(ctx context.Context, input GetUsersByIdsRequest) (GetUsersByIdsCompleteResult, error) {
	return c.GetUsersByIdsCompleteMatchingPredicate(ctx, input, DirectoryObjectOperationPredicate{})
}

// GetUsersByIdsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserClient) GetUsersByIdsCompleteMatchingPredicate(ctx context.Context, input GetUsersByIdsRequest, predicate DirectoryObjectOperationPredicate) (result GetUsersByIdsCompleteResult, err error) {
	items := make([]beta.DirectoryObject, 0)

	resp, err := c.GetUsersByIds(ctx, input)
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

	result = GetUsersByIdsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

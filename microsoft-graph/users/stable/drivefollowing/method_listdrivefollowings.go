package drivefollowing

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

type ListDriveFollowingsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DriveItem
}

type ListDriveFollowingsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DriveItem
}

type ListDriveFollowingsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListDriveFollowingsOperationOptions() ListDriveFollowingsOperationOptions {
	return ListDriveFollowingsOperationOptions{}
}

func (o ListDriveFollowingsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveFollowingsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListDriveFollowingsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveFollowingsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveFollowingsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveFollowings - Get following from users. The list of items the user is following. Only in OneDrive for
// Business.
func (c DriveFollowingClient) ListDriveFollowings(ctx context.Context, id stable.UserIdDriveId, options ListDriveFollowingsOperationOptions) (result ListDriveFollowingsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveFollowingsCustomPager{},
		Path:          fmt.Sprintf("%s/following", id.ID()),
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
		Values *[]stable.DriveItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveFollowingsComplete retrieves all the results into a single object
func (c DriveFollowingClient) ListDriveFollowingsComplete(ctx context.Context, id stable.UserIdDriveId, options ListDriveFollowingsOperationOptions) (ListDriveFollowingsCompleteResult, error) {
	return c.ListDriveFollowingsCompleteMatchingPredicate(ctx, id, options, DriveItemOperationPredicate{})
}

// ListDriveFollowingsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveFollowingClient) ListDriveFollowingsCompleteMatchingPredicate(ctx context.Context, id stable.UserIdDriveId, options ListDriveFollowingsOperationOptions, predicate DriveItemOperationPredicate) (result ListDriveFollowingsCompleteResult, err error) {
	items := make([]stable.DriveItem, 0)

	resp, err := c.ListDriveFollowings(ctx, id, options)
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

	result = ListDriveFollowingsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

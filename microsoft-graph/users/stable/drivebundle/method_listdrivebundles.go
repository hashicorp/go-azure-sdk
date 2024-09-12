package drivebundle

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

type ListDriveBundlesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DriveItem
}

type ListDriveBundlesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DriveItem
}

type ListDriveBundlesOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListDriveBundlesOperationOptions() ListDriveBundlesOperationOptions {
	return ListDriveBundlesOperationOptions{}
}

func (o ListDriveBundlesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveBundlesOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveBundlesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveBundlesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveBundlesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveBundles - Get bundles from users. Collection of bundles (albums and multi-select-shared sets of items). Only
// in personal OneDrive.
func (c DriveBundleClient) ListDriveBundles(ctx context.Context, id stable.UserIdDriveId, options ListDriveBundlesOperationOptions) (result ListDriveBundlesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveBundlesCustomPager{},
		Path:          fmt.Sprintf("%s/bundles", id.ID()),
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

// ListDriveBundlesComplete retrieves all the results into a single object
func (c DriveBundleClient) ListDriveBundlesComplete(ctx context.Context, id stable.UserIdDriveId, options ListDriveBundlesOperationOptions) (ListDriveBundlesCompleteResult, error) {
	return c.ListDriveBundlesCompleteMatchingPredicate(ctx, id, options, DriveItemOperationPredicate{})
}

// ListDriveBundlesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveBundleClient) ListDriveBundlesCompleteMatchingPredicate(ctx context.Context, id stable.UserIdDriveId, options ListDriveBundlesOperationOptions, predicate DriveItemOperationPredicate) (result ListDriveBundlesCompleteResult, err error) {
	items := make([]stable.DriveItem, 0)

	resp, err := c.ListDriveBundles(ctx, id, options)
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

	result = ListDriveBundlesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

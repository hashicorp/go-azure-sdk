package driverootchild

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

type ListDriveRootChildrenOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DriveItem
}

type ListDriveRootChildrenCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DriveItem
}

type ListDriveRootChildrenOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListDriveRootChildrenOperationOptions() ListDriveRootChildrenOperationOptions {
	return ListDriveRootChildrenOperationOptions{}
}

func (o ListDriveRootChildrenOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveRootChildrenOperationOptions) ToOData() *odata.Query {
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
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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

func (o ListDriveRootChildrenOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveRootChildrenCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveRootChildrenCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveRootChildren - Get children from users. Collection containing Item objects for the immediate children of
// Item. Only items representing folders have children. Read-only. Nullable.
func (c DriveRootChildClient) ListDriveRootChildren(ctx context.Context, id beta.UserIdDriveId, options ListDriveRootChildrenOperationOptions) (result ListDriveRootChildrenOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveRootChildrenCustomPager{},
		Path:          fmt.Sprintf("%s/root/children", id.ID()),
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
		Values *[]beta.DriveItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveRootChildrenComplete retrieves all the results into a single object
func (c DriveRootChildClient) ListDriveRootChildrenComplete(ctx context.Context, id beta.UserIdDriveId, options ListDriveRootChildrenOperationOptions) (ListDriveRootChildrenCompleteResult, error) {
	return c.ListDriveRootChildrenCompleteMatchingPredicate(ctx, id, options, DriveItemOperationPredicate{})
}

// ListDriveRootChildrenCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveRootChildClient) ListDriveRootChildrenCompleteMatchingPredicate(ctx context.Context, id beta.UserIdDriveId, options ListDriveRootChildrenOperationOptions, predicate DriveItemOperationPredicate) (result ListDriveRootChildrenCompleteResult, err error) {
	items := make([]beta.DriveItem, 0)

	resp, err := c.ListDriveRootChildren(ctx, id, options)
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

	result = ListDriveRootChildrenCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

package driveitemchild

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

type ListDriveItemChildrenOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.DriveItem
}

type ListDriveItemChildrenCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.DriveItem
}

type ListDriveItemChildrenOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListDriveItemChildrenOperationOptions() ListDriveItemChildrenOperationOptions {
	return ListDriveItemChildrenOperationOptions{}
}

func (o ListDriveItemChildrenOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveItemChildrenOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveItemChildrenOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveItemChildrenCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveItemChildrenCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveItemChildren - Get children from me. Collection containing Item objects for the immediate children of Item.
// Only items representing folders have children. Read-only. Nullable.
func (c DriveItemChildClient) ListDriveItemChildren(ctx context.Context, id stable.MeDriveIdItemId, options ListDriveItemChildrenOperationOptions) (result ListDriveItemChildrenOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveItemChildrenCustomPager{},
		Path:          fmt.Sprintf("%s/children", id.ID()),
		RetryFunc:     options.RetryFunc,
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

// ListDriveItemChildrenComplete retrieves all the results into a single object
func (c DriveItemChildClient) ListDriveItemChildrenComplete(ctx context.Context, id stable.MeDriveIdItemId, options ListDriveItemChildrenOperationOptions) (ListDriveItemChildrenCompleteResult, error) {
	return c.ListDriveItemChildrenCompleteMatchingPredicate(ctx, id, options, DriveItemOperationPredicate{})
}

// ListDriveItemChildrenCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveItemChildClient) ListDriveItemChildrenCompleteMatchingPredicate(ctx context.Context, id stable.MeDriveIdItemId, options ListDriveItemChildrenOperationOptions, predicate DriveItemOperationPredicate) (result ListDriveItemChildrenCompleteResult, err error) {
	items := make([]stable.DriveItem, 0)

	resp, err := c.ListDriveItemChildren(ctx, id, options)
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

	result = ListDriveItemChildrenCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

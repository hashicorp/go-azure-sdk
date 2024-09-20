package drivelistcontenttypecolumn

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

type ListDriveListContentTypeColumnsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ColumnDefinition
}

type ListDriveListContentTypeColumnsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ColumnDefinition
}

type ListDriveListContentTypeColumnsOperationOptions struct {
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

func DefaultListDriveListContentTypeColumnsOperationOptions() ListDriveListContentTypeColumnsOperationOptions {
	return ListDriveListContentTypeColumnsOperationOptions{}
}

func (o ListDriveListContentTypeColumnsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDriveListContentTypeColumnsOperationOptions) ToOData() *odata.Query {
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

func (o ListDriveListContentTypeColumnsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDriveListContentTypeColumnsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDriveListContentTypeColumnsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDriveListContentTypeColumns - Get columns from users. The collection of column definitions for this content type.
func (c DriveListContentTypeColumnClient) ListDriveListContentTypeColumns(ctx context.Context, id beta.UserIdDriveIdListContentTypeId, options ListDriveListContentTypeColumnsOperationOptions) (result ListDriveListContentTypeColumnsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDriveListContentTypeColumnsCustomPager{},
		Path:          fmt.Sprintf("%s/columns", id.ID()),
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
		Values *[]beta.ColumnDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDriveListContentTypeColumnsComplete retrieves all the results into a single object
func (c DriveListContentTypeColumnClient) ListDriveListContentTypeColumnsComplete(ctx context.Context, id beta.UserIdDriveIdListContentTypeId, options ListDriveListContentTypeColumnsOperationOptions) (ListDriveListContentTypeColumnsCompleteResult, error) {
	return c.ListDriveListContentTypeColumnsCompleteMatchingPredicate(ctx, id, options, ColumnDefinitionOperationPredicate{})
}

// ListDriveListContentTypeColumnsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DriveListContentTypeColumnClient) ListDriveListContentTypeColumnsCompleteMatchingPredicate(ctx context.Context, id beta.UserIdDriveIdListContentTypeId, options ListDriveListContentTypeColumnsOperationOptions, predicate ColumnDefinitionOperationPredicate) (result ListDriveListContentTypeColumnsCompleteResult, err error) {
	items := make([]beta.ColumnDefinition, 0)

	resp, err := c.ListDriveListContentTypeColumns(ctx, id, options)
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

	result = ListDriveListContentTypeColumnsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

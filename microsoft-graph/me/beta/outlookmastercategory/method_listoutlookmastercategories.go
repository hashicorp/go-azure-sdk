package outlookmastercategory

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

type ListOutlookMasterCategoriesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OutlookCategory
}

type ListOutlookMasterCategoriesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OutlookCategory
}

type ListOutlookMasterCategoriesOperationOptions struct {
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

func DefaultListOutlookMasterCategoriesOperationOptions() ListOutlookMasterCategoriesOperationOptions {
	return ListOutlookMasterCategoriesOperationOptions{}
}

func (o ListOutlookMasterCategoriesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOutlookMasterCategoriesOperationOptions) ToOData() *odata.Query {
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

func (o ListOutlookMasterCategoriesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListOutlookMasterCategoriesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOutlookMasterCategoriesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOutlookMasterCategories - List masterCategories. Get all the categories that have been defined for a user.
func (c OutlookMasterCategoryClient) ListOutlookMasterCategories(ctx context.Context, options ListOutlookMasterCategoriesOperationOptions) (result ListOutlookMasterCategoriesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListOutlookMasterCategoriesCustomPager{},
		Path:          "/me/outlook/masterCategories",
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
		Values *[]beta.OutlookCategory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOutlookMasterCategoriesComplete retrieves all the results into a single object
func (c OutlookMasterCategoryClient) ListOutlookMasterCategoriesComplete(ctx context.Context, options ListOutlookMasterCategoriesOperationOptions) (ListOutlookMasterCategoriesCompleteResult, error) {
	return c.ListOutlookMasterCategoriesCompleteMatchingPredicate(ctx, options, OutlookCategoryOperationPredicate{})
}

// ListOutlookMasterCategoriesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutlookMasterCategoryClient) ListOutlookMasterCategoriesCompleteMatchingPredicate(ctx context.Context, options ListOutlookMasterCategoriesOperationOptions, predicate OutlookCategoryOperationPredicate) (result ListOutlookMasterCategoriesCompleteResult, err error) {
	items := make([]beta.OutlookCategory, 0)

	resp, err := c.ListOutlookMasterCategories(ctx, options)
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

	result = ListOutlookMasterCategoriesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}

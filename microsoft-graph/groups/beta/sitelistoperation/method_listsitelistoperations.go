package sitelistoperation

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

type ListSiteListOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.RichLongRunningOperation
}

type ListSiteListOperationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.RichLongRunningOperation
}

type ListSiteListOperationsOperationOptions struct {
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

func DefaultListSiteListOperationsOperationOptions() ListSiteListOperationsOperationOptions {
	return ListSiteListOperationsOperationOptions{}
}

func (o ListSiteListOperationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSiteListOperationsOperationOptions) ToOData() *odata.Query {
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

func (o ListSiteListOperationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSiteListOperationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSiteListOperationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSiteListOperations - Get operations from groups. The collection of long-running operations on the list.
func (c SiteListOperationClient) ListSiteListOperations(ctx context.Context, id beta.GroupIdSiteIdListId, options ListSiteListOperationsOperationOptions) (result ListSiteListOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSiteListOperationsCustomPager{},
		Path:          fmt.Sprintf("%s/operations", id.ID()),
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
		Values *[]beta.RichLongRunningOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSiteListOperationsComplete retrieves all the results into a single object
func (c SiteListOperationClient) ListSiteListOperationsComplete(ctx context.Context, id beta.GroupIdSiteIdListId, options ListSiteListOperationsOperationOptions) (ListSiteListOperationsCompleteResult, error) {
	return c.ListSiteListOperationsCompleteMatchingPredicate(ctx, id, options, RichLongRunningOperationOperationPredicate{})
}

// ListSiteListOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SiteListOperationClient) ListSiteListOperationsCompleteMatchingPredicate(ctx context.Context, id beta.GroupIdSiteIdListId, options ListSiteListOperationsOperationOptions, predicate RichLongRunningOperationOperationPredicate) (result ListSiteListOperationsCompleteResult, err error) {
	items := make([]beta.RichLongRunningOperation, 0)

	resp, err := c.ListSiteListOperations(ctx, id, options)
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

	result = ListSiteListOperationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
